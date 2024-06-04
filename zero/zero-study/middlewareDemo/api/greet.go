package main

import (
	"flag"
	"fmt"

	"binggo/zero/zero-study/middlewareDemo/api/internal/config"
	"binggo/zero/zero-study/middlewareDemo/api/internal/handler"
	"binggo/zero/zero-study/middlewareDemo/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middlewareDemoFunc) // 在这加的是全局中间件

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
