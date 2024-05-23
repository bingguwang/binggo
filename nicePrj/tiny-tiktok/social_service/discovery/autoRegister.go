package discovery

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"tiny-tiktok/social_service/internal/handler"
	"tiny-tiktok/social_service/internal/service"
	"tiny-tiktok/utils/etcd"
)

// AutoRegister etcd自动注册
func AutoRegister() {
	etcdAddress := viper.GetString("etcd.address")
	etcdRegister, err := etcd.NewEtcdRegister(etcdAddress)

	if err != nil {
		log.Fatal(err)
	}

	serviceName := viper.GetString("server.name")
	serviceAddress := viper.GetString("server.address")
	err = etcdRegister.ServiceRegister(serviceName, serviceAddress, 30)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service.RegisterSocialServiceServer(server, handler.NewSocialService())

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
