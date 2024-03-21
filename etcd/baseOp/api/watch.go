package api

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

/**
 * @Description: 监控键值变化
 * @receiver c
 * @param key：键
 * @return clientv3.WatchChan：接收变化消息的channel
 */
func (c *EtcdClienter) Watch(key string) clientv3.WatchChan {
	rch := c.cli.Watch(context.Background(), key)
	return rch
}

/**
 * @Description: 监控具有某些前缀的键值变化
 * @receiver c
 * @param key：键
 * @return clientv3.WatchChan：接收变化消息的channel
 */
func (c *EtcdClienter) WatchWithPrefix(key string) clientv3.WatchChan {
	rch := c.cli.Watch(context.Background(), key, clientv3.WithPrefix())
	return rch
}

/**
 * @Description:
 * @receiver c
 * @param key：键
 * @return clientv3.WatchChan：接收变化消息的channel
 */
func (c *EtcdClienter) WatchWithProgressNotify(key string) clientv3.WatchChan {
	rch := c.cli.Watch(context.Background(), key, clientv3.WithProgressNotify())
	return rch
}
