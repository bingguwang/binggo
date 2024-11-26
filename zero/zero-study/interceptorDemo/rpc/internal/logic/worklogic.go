package logic

import (
	"context"
	"time"

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
	time.Sleep(200 * time.Millisecond) // 模拟处理时间，使限流器效果明显
	l.Info("调用rpcwork成功，work执行")

	return &bing.Response{
		Pong: "succ",
	}, nil
}
