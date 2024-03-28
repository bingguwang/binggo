package etcdDiscovery

import (
	regdis "binggo/etcd/registerDiscovery"
)

type RedisInfo struct {
	regdis.CommonInfo
	Name     string
	Password string
}

type RedisRegister struct {
	info    *RedisInfo
	service *regdis.ServiceRegister
}

func RegRedis(serviceName string, uuid string, regInfo *RedisInfo, timeOut int64) (*RedisRegister, error) {
	redisRegister := &RedisRegister{}

	serviceInfo := regdis.ServiceInfoType{
		regdis.HOST:   regInfo.Host,
		regdis.PORT:   regInfo.Port,
		regdis.NAME:   regInfo.Name,
		regdis.PASSWD: regInfo.Password,
	}
	serviceRegister, err := regdis.Register(serviceName, uuid, serviceInfo, timeOut)
	if err != nil {
		return redisRegister, err
	}

	redisRegister.info = regInfo
	redisRegister.service = serviceRegister

	return redisRegister, nil
}

func (c *RedisRegister) Update(regInfo *RedisInfo) error {
	serviceInfo := regdis.ServiceInfoType{
		regdis.HOST:   regInfo.Host,
		regdis.PORT:   regInfo.Port,
		regdis.NAME:   regInfo.Name,
		regdis.PASSWD: regInfo.Password,
	}
	return c.service.UpdateServiceInfo(serviceInfo)
}

func (c *RedisRegister) GetInfo() (*RedisInfo, error) {
	regInfo := &RedisInfo{}
	info, err := c.service.GetServiceInfo()
	if err != nil {
		return regInfo, err
	}
	if _, ok := info[regdis.HOST]; ok {
		regInfo.Host = info[regdis.HOST]
	}
	if _, ok := info[regdis.PORT]; ok {
		regInfo.Port = info[regdis.PORT]
	}
	if _, ok := info[regdis.NAME]; ok {
		regInfo.Name = info[regdis.NAME]
	}
	if _, ok := info[regdis.PASSWD]; ok {
		regInfo.Password = info[regdis.PASSWD]
	}

	return regInfo, nil
}

func (c *RedisRegister) UnRegister() error {
	return c.service.UnRegister()
}
