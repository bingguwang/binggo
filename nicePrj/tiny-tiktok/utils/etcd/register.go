package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
	"tiny-tiktok/user_service/pkg/logger"
)

// EtcdRegister 服务注册
type EtcdRegister struct {
	etcdCli            *clientv3.Client // etcdClient对象
	leaseId            clientv3.LeaseID // 租约id
	registerTime       string           // 注册时间
	serviceName        string           // 服务名称
	leaseKeepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
}

// CreateLease 创建租约。expire表示有效期(s)
func (e *EtcdRegister) CreateLease(expire int64) error {

	lease, err := e.etcdCli.Grant(context.Background(), expire)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	e.leaseId = lease.ID // 记录生成的租约Id
	return nil
}

// BindLease 绑定租约。将租约与对应的key-value绑定
func (e *EtcdRegister) BindLease(key string, value string) error {

	res, err := e.etcdCli.Put(context.Background(), key, value, clientv3.WithLease(e.leaseId))
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	logger.Log.Infof("bind lease success %v \n", res)
	return nil
}

// KeepAlive 获取续约通道 并 持续续租
func (e *EtcdRegister) KeepAlive() error {
	keepAliveChan, err := e.etcdCli.KeepAlive(context.Background(), e.leaseId)

	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	// 续约消息回复通道需要监听
	go func(keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse) {
		count := 0
		for {
			select {
			case resp := <-keepAliveChan:
				count++
				logger.Log.Infof("服务[%s]续约成功...leaseID=%d", e.serviceName, resp.ID)
			}
		}
		logger.Log.Infof("服务[%s]心跳停止, 服务需要重新注册了", e.serviceName)
	}(keepAliveChan)

	return nil
}

// Close 关闭服务
func (e *EtcdRegister) Close() error {
	logger.Log.Infof("close...\n")
	// 撤销租约
	e.etcdCli.Revoke(context.Background(), e.leaseId)
	return e.etcdCli.Close()
}

// NewEtcdRegister 初始化etcd服务注册对象
func NewEtcdRegister(etcdServerAddr, etcdPassword, etcdUsername string) (*EtcdRegister, error) {
	logger.Log.Info("start")

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdServerAddr},
		Password:    etcdPassword,
		Username:    etcdUsername,
		DialTimeout: 3 * time.Second,
	})

	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	logger.Log.Info("over")
	e := &EtcdRegister{
		etcdCli: client,
	}
	return e, nil
}

// ServiceRegister 服务注册。expire表示过期时间,serviceName和serviceAddr分别是服务名与服务地址
func (e *EtcdRegister) ServiceRegister(serviceName, serviceAddr string, expire int64) (err error) {

	// 创建租约
	err = e.CreateLease(expire)
	if err != nil {
		return err
	}
	logger.Log.Info("创建租约成功")
	// 将租约与k-v绑定
	err = e.BindLease(serviceName, serviceAddr)
	if err != nil {
		return err
	}
	logger.Log.Info("将租约与k-v绑定成功")

	// 持续续租
	err = e.KeepAlive()
	return err
}
