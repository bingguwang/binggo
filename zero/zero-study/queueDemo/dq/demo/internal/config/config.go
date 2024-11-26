package config

import (
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	DqConf struct {
		Beanstalks []dq.Beanstalk
		RedisConf  []redis.RedisConf
	}
}
