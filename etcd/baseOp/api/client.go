package api

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func ClientInit(dialTimeout time.Duration) (*EtcdClienter, error) {
	endpoints := []string{
		"192.168.2.44:2379",
	}
	if dialTimeout == 0 {
		dialTimeout = 5 * time.Second
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		return nil, err
	}

	return &EtcdClienter{
		cli:         cli,
		endpoints:   endpoints,
		dialTimeout: dialTimeout,
	}, nil
}
