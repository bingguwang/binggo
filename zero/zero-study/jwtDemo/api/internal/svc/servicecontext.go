package svc

import (
	"binggo/zero/zero-study/jwtDemo/api/internal/config"
	"binggo/zero/zero-study/jwtDemo/rpc/user/user"
	"binggo/zero/zero-study/jwtDemo/rpc/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserClient // 需要初始化实例化一次
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)), // 注意实例的不是NewUserClient，在它之上还有个封装:User接口
	}
}
