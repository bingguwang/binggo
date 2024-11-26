package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"user/api/internal/config"
	"user/log"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.SetWriter(*log.LogxKafka())
	return &ServiceContext{
		Config: c,
	}
}
