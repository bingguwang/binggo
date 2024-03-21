package api

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

type EtcdClienter struct {
	cli         *clientv3.Client
	endpoints   []string
	dialTimeout time.Duration
}

/**
 * @Description: 执行PUT操作，存入键值对
 * @receiver c
 * @param key：键
 * @param value：值
 * @param leaseID：lease ID，为0时表示不使用lease
 * @param requestTimeout：操作超时时间
 * @return *clientv3.PutResponse：PUT操作结果对象指针
 * @return error
 */
func (c *EtcdClienter) Put(key string, value string, leaseID clientv3.LeaseID, requestTimeout time.Duration) (*clientv3.PutResponse, error) {
	var err error
	var resp *clientv3.PutResponse

	if requestTimeout == 0 {
		requestTimeout = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	if leaseID > 0 {
		resp, err = c.cli.Put(ctx, key, value, clientv3.WithLease(leaseID))
	} else {
		resp, err = c.cli.Put(ctx, key, value)
	}
	if err != nil {
		switch err {
		case context.Canceled:
			log.Printf("[etcd] ctx is canceled by another routine: %v\n", err)
		case context.DeadlineExceeded:
			log.Printf("[etcd] ctx is attached with a deadline is exceeded: %v\n", err)
		case rpctypes.ErrEmptyKey:
			log.Printf("[etcd] client-side error: %v\n", err)
		default:
			log.Printf("[etcd] bad cluster endpoints, which are not etcd servers: %v\n", err)
		}
	}

	return resp, err
}

/**
 * @Description: 执行GET操作，获取键值
 * @receiver c
 * @param key：键
 * @param requestTimeout：操作超时时间
 * @return string：键对应的值
 * @return error
 */
func (c *EtcdClienter) Get(key string, requestTimeout time.Duration) (string, error) {
	if requestTimeout == 0 {
		requestTimeout = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	resp, err := c.cli.Get(ctx, key)
	if err != nil {
		return "", err
	}

	return string(resp.Kvs[0].Value), nil
}

/**
 * @Description: 执行GET操作，获取某个版本的键值
 * @receiver c
 * @param key：键
 * @param Revision：版本号
 * @param requestTimeout：操作超时时间
 * @return string：对应版本的键值
 * @return error
 */
func (c *EtcdClienter) GetWithRevision(key string, Revision int64, requestTimeout time.Duration) (string, error) {
	if requestTimeout == 0 {
		requestTimeout = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	resp, err := c.cli.Get(ctx, key, clientv3.WithRev(Revision))
	if err != nil {
		return "", err
	}

	return string(resp.Kvs[0].Value), nil
}

/**
 * @Description: 执行GET操作，获取有某些前缀的键值
 * @receiver c
 * @param key：键前缀
 * @param requestTimeout：操作超时时间
 * @return map[string]string：键值映射表
 * @return error
 */
func (c *EtcdClienter) GetSortedPrefix(key string, requestTimeout time.Duration) (map[string]string, error) {
	if requestTimeout == 0 {
		requestTimeout = 5 * time.Second
	}
	result := make(map[string]string, 0)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	resp, err := c.cli.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, err
	}

	for _, ev := range resp.Kvs {
		result[string(ev.Key)] = string(ev.Value)
	}

	return result, nil
}

/**
 * @Description: 执行DELETE操作，删除键值对
 * @receiver c
 * @param key：键
 * @param requestTimeout：操作超时时间
 * @return string：删除键对应的值
 * @return error
 */
func (c *EtcdClienter) Delete(key string, requestTimeout time.Duration) (string, error) {
	if requestTimeout == 0 {
		requestTimeout = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	resp, err := c.cli.Delete(ctx, key, clientv3.WithPrevKV())
	if err != nil {
		return "", err
	}

	return string(resp.PrevKvs[0].Value), nil
}

/**
 * @Description: 执行DELETE操作，删除具有某些前缀的键值对
 * @receiver c
 * @param key：键
 * @param requestTimeout：操作超时时间
 * @return map[string]string：删除的键值映射表
 * @return error
 */
func (c *EtcdClienter) DeletePrefix(key string, requestTimeout time.Duration) (map[string]string, error) {
	if requestTimeout == 0 {
		requestTimeout = 5 * time.Second
	}
	result := make(map[string]string, 0)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	resp, err := c.cli.Delete(ctx, key, clientv3.WithPrefix(), clientv3.WithPrevKV())
	if err != nil {
		return nil, err
	}

	for _, ev := range resp.PrevKvs {
		result[string(ev.Key)] = string(ev.Value)
	}

	return result, nil
}

/**
 * @Description: 压缩指定版本前的历史事件
 * @receiver c
 * @param Revision：版本号
 * @param requestTimeout：操作超时时间
 * @return error
 */
// TODO：未测试
func (c *EtcdClienter) Compact(Revision int64, requestTimeout time.Duration) error {
	if requestTimeout == 0 {
		requestTimeout = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	_, err := c.cli.Compact(ctx, Revision)

	if err != nil {
		return err
	}

	return nil
}

/**
 * @Description: 仅支持多个kv put操作的事务
 * @receiver c
 * @param kv：键值对
 * @param leaseID：租约ID，为0时表示不使用lease
 * @param requestTimeout：操作超时时间
 * @return error
 */
func (c *EtcdClienter) Txn(kv map[string]string, leaseID clientv3.LeaseID, requestTimeout time.Duration) error {
	if requestTimeout == 0 {
		requestTimeout = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	txn := c.cli.Txn(ctx)

	var ops []clientv3.Op
	for k, v := range kv {
		if leaseID > 0 {
			ops = append(ops, clientv3.OpPut(k, v, clientv3.WithLease(leaseID)))
		} else {
			ops = append(ops, clientv3.OpPut(k, v))
		}
	}

	_, err := txn.Then(ops...).Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *EtcdClienter) Ctx() context.Context {
	return c.cli.Ctx()
}

func (c *EtcdClienter) Close() error {
	return c.cli.Close()
}
