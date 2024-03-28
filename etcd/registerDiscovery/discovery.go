package register

import (
	api2 "binggo/etcd/baseOp/api"
	"binggo/redis"
	"context"
	"fmt"
	"strings"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

type ServiceDiscovery struct {
	cli         *api2.EtcdClienter
	ServiceName string
	ServerInfo  map[string]ServiceInfoType // uuid 唯一标识一个服务
	lock        sync.Mutex
	EventChan   chan map[string]string // 通知应用程序配置发生变化
}

/**
 * @Description: 服务发现
 * @param serviceName：要发现的服务名称
 * @return *ServiceDiscovery：
 * @return error
 */
func Discovery(serviceName string) (*ServiceDiscovery, error) {
	cli, err := api2.ClientInit(0)
	if err != nil {
		return nil, err
	}

	s := &ServiceDiscovery{
		cli:         cli,
		ServiceName: string(serviceName),
		ServerInfo:  make(map[string]ServiceInfoType, 0),
		EventChan:   make(chan map[string]string, 1),
	}

	// 根据前缀获取现有的key
	info, err := cli.GetSortedPrefix(fmt.Sprintf("%s/%s/", ServicePrefix, s.ServiceName), 0)
	if err != nil {
		s.Close()
		return nil, err
	}
	fmt.Println("info")
	fmt.Println(redis.TojsonStr(info))

	for k, v := range info {
		s.updateService(k, v) // 所有的key切割出uuid，作为key加入到serviceInfo里
	}

	// TODO: 在初始化和监听之间如果服务去注册了，则丢失了这段监听变化
	//  如果先起监听再获取数据则监听的数据可能被初始数据覆盖导致不准确
	go s.watcher(cli.Ctx())

	return s, nil
}

/**
 * @Description:监视服务变更
 * @receiver s
 */
func (s *ServiceDiscovery) watcher(ctx context.Context) {
	rch := s.cli.WatchWithPrefix(fmt.Sprintf("%s/%s/", ServicePrefix, s.ServiceName))
	for {
		select {
		case <-ctx.Done():
			//fmt.Println("cli ctx done, close watcher")
			return
		case wresp := <-rch:
			for _, ev := range wresp.Events {
				k := string(ev.Kv.Key)
				v := string(ev.Kv.Value)
				switch ev.Type {
				case mvccpb.PUT:
					//fmt.Println("PUT", k, v)
					s.updateService(k, v)
				case mvccpb.DELETE:
					//fmt.Println("DELETE", k)
					key := strings.Split(k, "/")
					if len(key) == 4 {
						uuid := key[3]
						s.delService(uuid)
					}
				}
			}
		}
	}
}

func (s *ServiceDiscovery) notify(info map[string]string) {
	select {
	case s.EventChan <- info:
		//fmt.Println(s.ServiceName, info)
	default:
		// 通道满时事件会丢失
		//fmt.Println("notify fail")
	}
}

/**
 * @Description: 更新服务信息
 * @receiver s
 * @param k：待更新服务信息的key，格式为etcd原生格式
 * @param v：待更新服务信息的value，格式为etcd原生格式
 */
func (s *ServiceDiscovery) updateService(k, v string) {
	key := strings.Split(k, "/")
	uuid := key[3]

	//fmt.Println(key, uuid, len(key))

	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.ServerInfo[uuid]; !ok {
		s.ServerInfo[uuid] = make(ServiceInfoType, 0)
	}
	if len(key) == 6 {
		subDir := key[4]
		if subDir == "config" {
			s.ServerInfo[uuid][key[len(key)-1]] = v
			// 当配置修改时，写入 <uuid> - <key> 的键值对，表示UUID下的key对应的value发生变化
			//s.notify(map[string]string{uuid: key[len(key)-1]})
			// 当配置删除时，写入 <put> - <uuid> 的键值对，表示uuid服务变更
			s.notify(map[string]string{"put": uuid})
		}
	}
}

/**
 * @Description: 删除已注销服务
 * @receiver s
 * @param uuid：注销服务对应的UUID
 */
func (s *ServiceDiscovery) delService(uuid string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.ServerInfo[uuid]; ok {
		delete(s.ServerInfo, uuid)
		// 当配置删除时，写入 <delete> - <uuid> 的键值对，表示uuid服务注销
		s.notify(map[string]string{"delete": uuid})
	}
}

/**
 * @Description: 获取服务信息
 * @receiver s
 * @return map[string]ServiceInfoType：UUID为键值
 */
func (s *ServiceDiscovery) GetServiceInfo() map[string]ServiceInfoType {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.ServerInfo
}

/**
 * @Description:关闭发现服务
 * @receiver s
 * @return error
 */
func (s *ServiceDiscovery) Close() error {
	close(s.EventChan)
	return s.cli.Close()
}
