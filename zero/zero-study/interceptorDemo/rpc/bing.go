package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/time/rate"
	"google.golang.org/grpc/status"

	"binggo/zero/zero-study/interceptorDemo/rpc/bing"
	"binggo/zero/zero-study/interceptorDemo/rpc/internal/config"
	"binggo/zero/zero-study/interceptorDemo/rpc/internal/server"
	"binggo/zero/zero-study/interceptorDemo/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/bing.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		bing.RegisterBingServer(grpcServer, server.NewBingServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(rateLimitInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

var limiter = rate.NewLimiter(rate.Limit(100), 100) // 实例化一个令牌桶拦截器，应该是全局共享的才有限流效果

// 这里自定义一个用于限流的服务端拦截器
func rateLimitInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if !limiter.Allow() {
		logx.Error("限流了")
		return nil, status.Error(500, "限流了")
	}
	logx.Info("正常处理rpc请求")
	return handler(ctx, req) // 处理rpc调用请求
}
