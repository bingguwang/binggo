package etcdDiscovery

import regDis "binggo/etcd/registerDiscovery"

type RedisDiscoveryer struct {
	service *regDis.ServiceDiscovery
}

func DiscoveryRedis(serviceName string) (*RedisDiscoveryer, error) {
	redisDiscoveryer := &RedisDiscoveryer{}
	serviceDiscovery, err := regDis.Discovery(serviceName)
	if err != nil {
		return redisDiscoveryer, err
	}

	redisDiscoveryer.service = serviceDiscovery

	return redisDiscoveryer, err
}

func (c *RedisDiscoveryer) Event() <-chan map[string]string {
	return c.service.EventChan
}

/**
 * @Description: 获取服务信息
 * @receiver s
 * @return map[string]RedisInfo：UUID为键值
 */
func (c *RedisDiscoveryer) GetServiceInfo() map[string]RedisInfo {
	serviceInfo := c.service.GetServiceInfo()

	redisInfo := make(map[string]RedisInfo, 0)
	for uuid, info := range serviceInfo {
		data := RedisInfo{}
		for k, v := range info {
			if k == "host" || k == "ip" {
				data.Host = v
			} else if k == "port" {
				data.Port = v
			} else if k == "name" {
				data.Name = v
			} else if k == "passwd" {
				data.Password = v
			}
		}
		redisInfo[uuid] = data
	}

	return redisInfo
}

func (c *RedisDiscoveryer) Close() error {
	return c.service.Close()
}
