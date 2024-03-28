package api

import (
	register "binggo/etcd/registerDiscovery"
	ed "binggo/redis/etcdDiscovery"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net"
)

type RedisProxyer struct {
	Pool          *redis.ClusterClient
	MinIdleConns  int
	SentinelAddrs []string
}

// 释放代理对象
func (p *RedisProxyer) Close() error {
	return p.Pool.Close()
}

/**
 * @Description: 新建缓存代理
 * @param minIdleConns：最小连接数，小于1时自动修正为1
 * @return *RedisProxyer
 * @return error
 */
func NewRedisProxy(minIdleConns int) (*RedisProxyer, error) {
	if minIdleConns < 1 {
		minIdleConns = 1
	}

	// 服务配置自动发现
	serviceDiscovery, err := ed.DiscoveryRedis(register.RedisServiceName)
	if err != nil {
		return nil, err
	}
	defer serviceDiscovery.Close()

	masterName := ""
	passwd := ""
	serviceInfo := serviceDiscovery.GetServiceInfo()
	sentinelAddrs := make([]string, 0)
	//fmt.Println(serviceInfo)
	if len(serviceInfo) == 0 {
		return nil, errors.New("no redis service available")
	}
	for _, info := range serviceInfo {
		// 默认只有一个redis数据库服务
		masterName = info.Name
		passwd = info.Password
		// 检查IP格式
		address := net.ParseIP(info.Host)
		if address == nil {
			return nil, errors.New("ip format error")
		}
		sentinelAddrs = append(sentinelAddrs, fmt.Sprintf("%s:%s", info.Host, info.Port))
	}
	fmt.Println(masterName, passwd, sentinelAddrs)

	pool, err := NewRedisPool(masterName, passwd, sentinelAddrs, minIdleConns)
	if err != nil {
		return nil, err
	}

	return &RedisProxyer{
		Pool:          pool,
		MinIdleConns:  minIdleConns,
		SentinelAddrs: sentinelAddrs,
	}, nil
}
