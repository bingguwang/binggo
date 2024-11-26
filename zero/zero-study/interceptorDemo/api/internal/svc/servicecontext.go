package svc

import (
	"binggo/zero/zero-study/interceptorDemo/api/internal/config"
	"binggo/zero/zero-study/interceptorDemo/rpc/bingclient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Bing   bingclient.Bing
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Bing:   bingclient.NewBing(zrpc.MustNewClient(c.Bing, zrpc.WithUnaryClientInterceptor(myTimeInterceptor))),
	}
}

// 这里客户端拦截器打印调用耗时时间
func myTimeInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	stime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...) // 执行rpc调用
	if err != nil {
		return err
	}

	logx.Infof("[client intercept]调用 %s 方法 耗时: %v\n", method, time.Now().Sub(stime))
	return nil
}
