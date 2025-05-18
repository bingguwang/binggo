package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ghettovoice/gosip"
	"github.com/ghettovoice/gosip/log"
)

var (
	logger log.Logger
)

func init() {
	logger = log.NewDefaultLogrusLogger().WithPrefix("Server")
}

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	srvConf := gosip.ServerConfig{} // sip服务器配置
	srv := gosip.NewServer(srvConf, nil, nil, logger)
	// 加密认证
	//srv.Listen("wss", "0.0.0.0:5081", &transport.TLSConfig{Cert: "certs/cert.pem", Key: "certs/key.pem"})

	srv.Listen("tcp", "0.0.0.0:5081")

	<-stop

	srv.Shutdown()
}
