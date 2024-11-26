package svc

import (
	"binggo/zero/zero-study/queueDemo/dq/demo/internal/config"
	"github.com/zeromicro/go-queue/dq"
)

type ServiceContext struct {
	Config         config.Config
	DqPusherClient dq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		DqPusherClient: dq.NewProducer(c.DqConf.Beanstalks),
		// 消费这也是类似，在这里实例化消费者
	}
}
