// 服务发现,发现所有的服务，返回一个map

package discovery

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tiny-tiktok/api_router/internal/service"
	"tiny-tiktok/api_router/pkg/logger"
	"tiny-tiktok/utils/etcd"
)

func Resolver() map[string]interface{} {
	serveInstance := make(map[string]interface{})

	etcdAddress := viper.GetString("etcd.address")
	etcdPassword := viper.GetString("etcd.password")
	etcdUsername := viper.GetString("etcd.username")
	serviceDiscovery, err := etcd.NewServiceDiscovery([]string{etcdAddress}, etcdPassword, etcdUsername)
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer serviceDiscovery.Close()

	// 获取用户服务实例
	err = serviceDiscovery.ServiceDiscovery("user_service")
	if err != nil {
		logger.Log.Fatal(err)
	}
	userServiceAddr, _ := serviceDiscovery.GetService("user_service")
	// 使用了容器部署与容器编排后，其实可以不用etcd注册中心了，因为注册中心无非就是存服务和地址的映射
	// 有容器编排之后，容器的hostname和port已经知道了的，所以无需注册中心其实
	userConn, err := grpc.Dial(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*50), // 设置最大接收消息大小为50MB
			grpc.MaxCallSendMsgSize(1024*1024*50), // 设置最大发送消息大小为50MB
		),
	)
	if err != nil {
		logger.Log.Fatal(err)
	}
	userClient := service.NewUserServiceClient(userConn)
	logger.Log.Info("获取用户服务实例--成功--")
	serveInstance["user_service"] = userClient

	// 获取视频服务实例
	err = serviceDiscovery.ServiceDiscovery("video_service")
	if err != nil {
		logger.Log.Fatal(err)
	}
	videoServiceAddr, _ := serviceDiscovery.GetService("video_service")
	videoConn, err := grpc.Dial(videoServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal(err)
	}

	videoClient := service.NewVideoServiceClient(videoConn)
	logger.Log.Info("获取视频服务实例--成功--")
	serveInstance["video_service"] = videoClient

	// 获取社交服务实例
	err = serviceDiscovery.ServiceDiscovery("social_service")
	if err != nil {
		logger.Log.Fatal(err)
	}
	socialServiceAddr, _ := serviceDiscovery.GetService("social_service")
	socialConn, err := grpc.Dial(socialServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal(err)
	}

	socialClient := service.NewSocialServiceClient(socialConn)
	logger.Log.Info("获取社交服务实例--成功--")
	serveInstance["social_service"] = socialClient

	//wrapper.NewWrapper("user_service")
	//wrapper.NewWrapper("video_service")
	//wrapper.NewWrapper("social_service")

	return serveInstance
}
