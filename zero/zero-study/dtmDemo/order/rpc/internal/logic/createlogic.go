package logic

import (
	"binggo/zero/zero-study/dtmDemo/order/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"

	"binggo/zero/zero-study/dtmDemo/order/rpc/internal/svc"
	"binggo/zero/zero-study/dtmDemo/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	l.Info("rpc调用order.Create 创建订单")
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		newOrder := model.Order{
			Uid:    uint64(in.Uid),
			Pid:    uint64(in.Pid),
			Amount: uint64(in.Amount),
			Status: 0,
		}
		// 创建订单
		_, err = l.svcCtx.OrderModel.TxInsert(l.ctx, tx, &newOrder)
		if err != nil {
			//if err == nil { // 测试一下子事务执行失败时发送的情况
			return fmt.Errorf("订单创建失败")
		}

		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.CreateResponse{}, nil
}
