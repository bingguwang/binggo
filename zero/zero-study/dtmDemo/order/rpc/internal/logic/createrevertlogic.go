package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"

	"binggo/zero/zero-study/dtmDemo/order/rpc/internal/svc"
	"binggo/zero/zero-study/dtmDemo/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRevertLogic {
	return &CreateRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateRevert 调用订单创建事务的补偿函数
func (l *CreateRevertLogic) CreateRevert(in *order.CreateRequest) (*order.CreateResponse, error) {
	l.Info("订单创建补偿操作:CreateRevert")
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	// 子事务屏障用于防止重复执行和确保幂等性。
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {

		l.Info("修改订单状态为失效(此过程里会有其他的一些rpc调用等等)")
		l.Info("修改订单状态9，标识订单已失效，并更新订单")

		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.CreateResponse{}, nil
}
