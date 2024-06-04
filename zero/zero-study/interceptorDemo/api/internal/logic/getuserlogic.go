package logic

import (
	"binggo/zero/zero-study/interceptorDemo/api/internal/svc"
	"binggo/zero/zero-study/interceptorDemo/api/internal/types"
	"binggo/zero/zero-study/interceptorDemo/rpc/bingclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.Request) (resp *types.Response, err error) {
	rpcResp, err := l.svcCtx.Bing.Work(l.ctx, &bingclient.Request{})
	return &types.Response{
		Msg: rpcResp.Pong,
	}
}
