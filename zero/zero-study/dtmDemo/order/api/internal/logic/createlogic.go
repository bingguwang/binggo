package logic

import (
	"binggo/zero/zero-study/dtmDemo/order/rpc/order"
	"binggo/zero/zero-study/dtmDemo/product/rpc/product"
	"context"
	"google.golang.org/grpc/status"

	"binggo/zero/zero-study/dtmDemo/order/api/internal/svc"
	"binggo/zero/zero-study/dtmDemo/order/api/internal/types"
	"github.com/dtm-labs/dtmgrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	l.Info("api调用Create...")
	// 获取 OrderRpc BuildTarget
	orderRpcBussiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()

	// 获取 ProductRpc BuildTarget
	productRpcBussiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}
	//orderRpcBussiServer 和 productRpcBussiServer 保存了 Order 和 Product 服务的 RPC 端点地址

	logx.Info("Order服务的 RPC 端点地址----", orderRpcBussiServer)
	logx.Info("Product服务的 RPC 端点地址--------", productRpcBussiServer)

	//注册 dtm 服务的 etcd 地址
	var dtmServer = "etcd://root:123456@192.168.2.130:2379/dtmservice"
	// 创建一个gid,  生成一个全局事务 ID（gid）
	gid := dtmgrpc.MustGenGid(dtmServer)
	// 创建一个saga协议的事务
	l.Info("创建一个saga协议的事务")
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		// 第一个参数是要在事务里执行的操作
		// 第二个参数是事务执行失败时的补偿操作
		// 建订单
		Add(orderRpcBussiServer+"/pay.Order/Create", orderRpcBussiServer+"/pay.Order/CreateRevert", &order.CreateRequest{
			Uid:    req.Uid,
			Pid:    req.Pid,
			Amount: req.Amount,
			Status: 0,
		}).
		// 第一个参数是要在事务里执行的操作
		// 第二个参数是事务执行失败时的补偿操作
		// 减库存
		Add(productRpcBussiServer+"/product.Product/DecrStock", productRpcBussiServer+"/product.Product/DecrStockRevert", &product.DecrStockRequest{
			Id:  req.Pid,
			Num: 1,
		})

	// 事务提交
	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateResponse{}, nil
}
