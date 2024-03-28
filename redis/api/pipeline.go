package api

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func (p *RedisProxyer) Pipeline(args [][]interface{}) ([]redis.Cmder, error) {
	pipe := p.Pool.Pipeline()
	defer pipe.Close()

	ctx := context.Background()
	for _, arg := range args {
		pipe.Do(ctx, arg...)
	}
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return cmds, err
}
