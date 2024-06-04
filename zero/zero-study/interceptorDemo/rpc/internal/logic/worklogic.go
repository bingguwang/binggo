package logic

import (
	"context"

	"binggo/zero/zero-study/interceptorDemo/rpc/bing"
	"binggo/zero/zero-study/interceptorDemo/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWorkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkLogic {
	return &WorkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WorkLogic) Work(in *bing.Request) (*bing.Response, error) {
	// todo: add your logic here and delete this line

	return &bing.Response{}, nil
}
