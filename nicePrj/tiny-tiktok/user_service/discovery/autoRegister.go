package discovery

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"time"
	"tiny-tiktok/user_service/internal/handler"
	"tiny-tiktok/user_service/internal/service"
	log "tiny-tiktok/user_service/pkg/logger"
	"tiny-tiktok/utils/etcd"
)

// AutoRegister etcd自动注册
func AutoRegister() {
	log.Log.Info("开始etcd自动注册")
	etcdAddress := viper.GetString("etcd.address")
	etcdPassword := viper.GetString("etcd.password")
	etcdUsername := viper.GetString("etcd.username")
	log.Log.Info(etcdAddress)
	log.Log.Info(etcdPassword)

	etcdRegister, err := etcd.NewEtcdRegister(etcdAddress, etcdPassword, etcdUsername)

	if err != nil {
		panic(err.Error())
	}

	serviceName := viper.GetString("server.name")
	serviceAddress := viper.GetString("server.address")
	log.Log.Infof("注册服务:name:%s,address:%s", serviceName, serviceAddress)
	err = etcdRegister.ServiceRegister(serviceName, serviceAddress, 30*int64(time.Minute.Seconds()))
	if err != nil {
		panic(err.Error())
	}

	log.Log.Infof("服务[%s]注册etcd成功", serviceName)
	listener, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		panic(err.Error())
	}

	serverOptions := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 50), // 设置最大接收消息大小为50MB
		grpc.MaxSendMsgSize(1024 * 1024 * 50), // 设置最大发送消息大小为50MB
	}
	server := grpc.NewServer(serverOptions...)
	service.RegisterUserServiceServer(server, handler.NewUserService())

	err = server.Serve(listener)
	if err != nil {
		panic(err.Error())
	}
}
