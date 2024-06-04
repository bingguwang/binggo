package svc

import (
	"binggo/zero/zero-study/interceptorDemo/api/internal/config"
	"binggo/zero/zero-study/interceptorDemo/rpc/bingclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Bing   bingclient.Bing
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Bing:   bingclient.NewBing(zrpc.MustNewClient(c.Bing)),
	}
}
