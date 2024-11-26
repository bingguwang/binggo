package svc

import (
	"binggo/zero/zero-study/dtmDemo/order/api/internal/config"
	"binggo/zero/zero-study/dtmDemo/order/rpc/orderclient"
	"binggo/zero/zero-study/dtmDemo/product/rpc/productclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRpc   orderclient.Order
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
