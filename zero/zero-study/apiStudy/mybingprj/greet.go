package main

import (
	"apiStudy/mybingprj/internal/config"
	"apiStudy/mybingprj/internal/handler"
	"apiStudy/mybingprj/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c) // 加载指定路径的配置文件

	server := rest.MustNewServer(c.RestConf) // 根据传入的配置返回一个rest server
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
