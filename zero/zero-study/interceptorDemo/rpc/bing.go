package main

import (
	"flag"
	"fmt"

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

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
