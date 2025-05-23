// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	orderFieldNames          = builder.RawFieldNames(&Order{})
	orderRows                = strings.Join(orderFieldNames, ",")
	orderRowsExpectAutoSet   = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderRowsWithPlaceHolder = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheOrderIdPrefix = "cache:pay:id:"
)

type (
	orderModel interface {
		Insert(ctx context.Context, data *Order) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Order, error)
		Update(ctx context.Context, data *Order) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultOrderModel struct {
		sqlc.CachedConn
		table string
	}

	Order struct {
		Id         uint64    `db:"id"`
		Uid        uint64    `db:"uid"`    // 用户ID
		Pid        uint64    `db:"pid"`    // 产品ID
		Amount     uint64    `db:"amount"` // 订单金额
		Status     uint64    `db:"status"` // 订单状态
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultOrderModel {
	return &defaultOrderModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`pay`",
	}
}

func (m *defaultOrderModel) Delete(ctx context.Context, id uint64) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) FindOne(ctx context.Context, id uint64) (*Order, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	var resp Order
	err := m.QueryRowCtx(ctx, &resp, orderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderModel) Insert(ctx context.Context, data *Order) (sql.Result, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uid, data.Pid, data.Amount, data.Status)
	}, orderIdKey)
	return ret, err
}

func (m *defaultOrderModel) Update(ctx context.Context, data *Order) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Pid, data.Amount, data.Status, data.Id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheOrderIdPrefix, primary)
}

func (m *defaultOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrderModel) tableName() string {
	return m.table
}
