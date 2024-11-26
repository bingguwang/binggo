package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		TxInsert(ctx context.Context, tx *sql.Tx, data *Order) (sql.Result, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c, opts...),
	}
}

// TxInsert 乍看上去多此一举？明明只有一个insert操作，为什么还要放在一个事务里？？
// 因为在分布式事务中，每个参与的服务或数据库节点都需要在事务中执行操作，以确保在全局事务提交或回滚时能够协调一致。
// 所以这里必须放在事务里
func (m *defaultOrderModel) TxInsert(ctx context.Context, tx *sql.Tx, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
	ret, err := tx.ExecContext(ctx, query, data.Uid, data.Pid, data.Amount, data.Status)

	return ret, err
}
