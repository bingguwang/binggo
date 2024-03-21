package api

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

/**
租约是个对象，会持久化的对象
一个租约对象可以绑定多个key
一个key只能绑定到一个租约对象
绑定使用PUT方法
*/

/**
 * @Description: 创建租约
 * @receiver c
 * @param ttl：过期时间，单位为秒
 * @return clientv3.LeaseID：租约ID
 * @return error
 */
func (c *EtcdClienter) Grant(ttl int64) (clientv3.LeaseID, error) {
	resp, err := c.cli.Grant(context.TODO(), ttl)
	if err != nil {
		return 0, err
	}
	return resp.ID, nil
}

/**
 * @Description: 撤销租约
 * @receiver c
 * @param leaseID：要操作的租约ID
 * @return error
 */
func (c *EtcdClienter) Revoke(leaseID clientv3.LeaseID) error {
	// revoking lease expires the key attached to its lease ID
	_, err := c.cli.Revoke(context.TODO(), leaseID)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @Description: 续约
 * @receiver c
 * @param leaseID：要操作的租约ID
 * @return <-chan：通道用来接收续约的回复消息
 * @return error
 */
func (c *EtcdClienter) KeepAlive(leaseID clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	ch, err := c.cli.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

/**
 * @Description: 续租一次
 * @receiver c
 * @param leaseID：要操作的租约ID
 * @return *clientv3.LeaseKeepAliveResponse
 * @return error
 */
func (c *EtcdClienter) KeepAliveOnce(leaseID clientv3.LeaseID) (*clientv3.LeaseKeepAliveResponse, error) {
	// to renew the lease only once
	ka, err := c.cli.KeepAliveOnce(context.TODO(), leaseID)
	if err != nil {
		return nil, err
	}

	return ka, nil
}
