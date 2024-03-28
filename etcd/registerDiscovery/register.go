package register

import (
	api2 "binggo/etcd/baseOp/api"
	"binggo/utils"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	// 服务前缀
	ServicePrefix = "/service"

	// 节点前缀
	NodePrefix = "/node"

	ONLINE  = "online"
	OFFLINE = "offline"

	// 服务字段
	HOST   = "host"
	IP     = "ip"
	PORT   = "port"
	NAME   = "name"
	USER   = "user"
	PASSWD = "passwd"
)

const (
	RedisServiceName = "redis_service"
)

var exitFlag = false

type ServiceInfoType map[string]string

// 服务注册
type ServiceRegister struct {
	cli           *api2.EtcdClienter                      // etcd客户端操作对象
	ttl           int64                                   // 租约超时时间
	leaseID       clientv3.LeaseID                        // 租约ID
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse // 租约回复接收通道
	ServiceName   string                                  // 服务名称
	UUID          string                                  // 用来区分不同进程启动同一服务
	RegisterTime  string                                  // 服务注册时间
	serviceInfo   ServiceInfoType
}

/**
 * @Description: 服务注册
 * @param serviceName：服务名称
 * @param uuid：服务唯一的标识，一个服务进程重启后UUID不能变化，不同服务UUID不能冲突
 * @param serviceInfo：服务信息
 * @param timeOut：租约超时时间即服务离线超时时间，正整数，单位为秒
 * @return *ServiceRegister：
 * @return error
 */
func Register(serviceName string, uuid string, serviceInfo ServiceInfoType, timeOut int64) (*ServiceRegister, error) {
	if timeOut < 0 {
		return nil, errors.New("timeOut should be greater than 0")
	}

	cli, err := api2.ClientInit(0)
	if err != nil {
		return nil, err
	}

	s := &ServiceRegister{
		cli:           cli,
		ttl:           timeOut,
		leaseID:       clientv3.NoLease,
		keepAliveChan: nil,
		ServiceName:   string(serviceName),
		UUID:          uuid,
		RegisterTime:  time.Now().Format("2006-01-02 15:04:05"),
		serviceInfo:   make(ServiceInfoType),
	}

	for k, v := range serviceInfo {
		s.serviceInfo[k] = v
	}

	fmt.Println("开始注册")
	err = s.reg() //注册
	if err != nil {
		_ = cli.Close()
		return nil, err
	}
	fmt.Println("注册完成")

	return s, nil
}

/**
 * @Description: 及时从通道取出续约回复，否则通道会被塞满弹出告警
 * @receiver s
 */
func (s *ServiceRegister) listenLeaseRespChan() {
	log.Println("[register] service", s.ServiceName, "heartbeat start")
	count := 0
	for leaseKeepResp := range s.keepAliveChan {
		_ = leaseKeepResp
		count++
		fmt.Println(s.ServiceName, "续约成功", leaseKeepResp, "count:", count)
	}
	fmt.Println("走到这说明心跳停止")
	log.Println("[register] service", s.ServiceName, "heartbeat stop")

	s.leaseID = clientv3.NoLease

	if !exitFlag {
		fmt.Println("重新注册")
		go s.reReg()
	}
}

func (s *ServiceRegister) reReg() {
	fmt.Println("reReg")

	for {
		if s.leaseID != clientv3.NoLease {
			return
		}

		log.Println("[register] try register again")
		err := s.reg()
		if err == nil {
			return
		}

		log.Println("[register] rereg error:", err.Error())
		time.Sleep(time.Second * 5)
	}
}

/**
 * @Description: 当租约异常时，提供重新注册机制
 * @receiver s
 * @return error
 */
func (s *ServiceRegister) reg() (err error) {
	if s.leaseID != clientv3.NoLease {
		return errors.New("the lease is normal")
	}

	var leaseRespChan <-chan *clientv3.LeaseKeepAliveResponse
	var nodeIp string
	metaInfo := make(map[string]string, 0)

	fmt.Println("设置租约时间")
	// 设置租约时间
	leaseID, err := s.cli.Grant(s.ttl)
	if err != nil {
		return err
	}
	s.leaseID = leaseID
	registerTime := time.Now().Format("2006-01-02 15:04:05")
	s.RegisterTime = registerTime

	// 注册服务并绑定租约,服务的key格式为 /service/服务名称/UUID
	nodeIp, err = utils.GetOutBoundIP()
	if err != nil {
		goto cleanLease
	}
	_, err = s.cli.Put(fmt.Sprintf("%s/%s/%s", ServicePrefix, s.ServiceName, s.UUID), nodeIp, s.leaseID, 0)
	if err != nil {
		goto cleanLease
	}

	// 获取元信息，包括但不限于：
	// register_time: 注册时间
	metaInfo["register_time"] = registerTime

	err = s.update(s.serviceInfo, "config")
	if err != nil {
		goto cleanLease
	}

	err = s.UpdateMetaInfo(metaInfo)
	if err != nil {
		goto cleanLease
	}

	// 保持心跳
	leaseRespChan, err = s.cli.KeepAlive(s.leaseID)
	if err != nil {
		goto cleanLease
	}
	s.keepAliveChan = leaseRespChan

	// 监听 keepAliveChan，查看心跳 也即查看租约续约情况
	go s.listenLeaseRespChan()

	return nil

cleanLease:
	s.cli.Revoke(s.leaseID)
	s.leaseID = clientv3.NoLease
	return err
}

func (s *ServiceRegister) UpdateServiceInfo(info ServiceInfoType) error {
	// 只更新变化的键值对
	updateInfo := make(map[string]string)

	for k, v := range info {
		_, ok := s.serviceInfo[k]
		if ok {
			if s.serviceInfo[k] != info[k] {
				updateInfo[k] = v
			}
		} else {
			updateInfo[k] = v
		}
	}

	err := s.update(updateInfo, "config")
	if err != nil {
		return err
	}

	s.serviceInfo = info
	return nil
}

func (s *ServiceRegister) UpdateMetaInfo(info ServiceInfoType) error {
	return s.update(info, "meta")
}

/*
*
  - @Description: 更新服务信息:
    不处理服务配置减少的情况
    服务本身键值的key格式为: /service/服务名称/UUID/config/key
    元数据键值的key格式为: /service/服务名称/UUID/meta/key
  - @receiver s
  - @param serviceInfo：服务的配置信息等
  - @param subDir: 可选值[config|meta]，分别用于[更新服务配置信息|更新集群管理自身需要元数据]
  - @return error
*/
func (s *ServiceRegister) update(serviceInfo ServiceInfoType, subDir string) error {
	if s.leaseID == clientv3.NoLease {
		return errors.New("re-registering, please try again later")
	}

	if subDir != "config" && subDir != "meta" {
		return errors.New("subDir should be one of [config|meta]")
	}

	kvs := make(map[string]string)
	for k, v := range serviceInfo {
		// 检查键里不能带'/'
		if strings.Contains(k, "/") {
			return errors.New("key cannot contain '/'")
		}

		kvs[fmt.Sprintf("%s/%s/%s/%s/%s", ServicePrefix, s.ServiceName, s.UUID, subDir, k)] = v
	}

	err := s.cli.Txn(kvs, s.leaseID, 0)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceRegister) GetServiceInfo() (ServiceInfoType, error) {
	return s.getInfo("config")
}

func (s *ServiceRegister) GetMetaInfo() (ServiceInfoType, error) {
	return s.getInfo("meta")
}

func (s *ServiceRegister) getInfo(subDir string) (ServiceInfoType, error) {
	serviceInfo := make(map[string]string, 0)

	info, err := s.cli.GetSortedPrefix(fmt.Sprintf("%s/%s/%s/%s/", ServicePrefix, s.ServiceName, s.UUID, subDir), 0)
	if err != nil {
		return nil, err
	}

	for k, v := range info {
		key := strings.Split(k, "/")
		if len(key) > 0 {
			serviceInfo[key[len(key)-1]] = v
		}
	}

	serviceInfo["serviceName"] = s.ServiceName

	return serviceInfo, nil
}

/**
 * @Description: 注销服务
 * @receiver s
 * @return error
 */
func (s *ServiceRegister) UnRegister() error {
	exitFlag = true
	if s.leaseID != clientv3.NoLease {
		if err := s.cli.Revoke(s.leaseID); err != nil {
			return err
		}
	}

	return s.cli.Close()
}
