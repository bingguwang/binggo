package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShorturlModel = (*customShorturlModel)(nil)

type (
	// ShorturlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShorturlModel.
	ShorturlModel interface {
		shorturlModel
		// 这里可以添加更多操作数据库的方法，方法格式可以参考shorturlModel

		//FindAll(ctx context.Context) ([]*Shorturl, error)
	}

	// 自定义的内容由此结构来完成,上面自定义的方法需要由此customShorturlModel结构来实现
	customShorturlModel struct {
		*defaultShorturlModel
	}
)

// NewShorturlModel returns a model for the database table.
func NewShorturlModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ShorturlModel {
	return &customShorturlModel{
		defaultShorturlModel: newShorturlModel(conn, c, opts...),
	}
}
