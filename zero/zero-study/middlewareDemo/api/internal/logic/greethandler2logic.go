package logic

import (
	"context"

	"binggo/zero/zero-study/middlewareDemo/api/internal/svc"
	"binggo/zero/zero-study/middlewareDemo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetHandler2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetHandler2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetHandler2Logic {
	return &GreetHandler2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetHandler2Logic) GreetHandler2(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	l.Info("GreetHandler2处理Greet请求...")
	return
}
