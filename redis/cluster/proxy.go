package cluster

import "github.com/go-redis/redis/v8"

type RedisProxyer struct {
	pool          *redis.ClusterClient
	minIdleConns  int
	sentinelAddrs []string
}
