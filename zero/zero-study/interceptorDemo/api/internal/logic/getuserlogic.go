package logic

import (
	"binggo/zero/zero-study/interceptorDemo/api/internal/svc"
	"binggo/zero/zero-study/interceptorDemo/api/internal/types"
	"binggo/zero/zero-study/interceptorDemo/rpc/bing"
	"binggo/zero/zero-study/interceptorDemo/rpc/bingclient"
	"context"
	"sync"

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
	var rpcResp *bing.Response
	var wg sync.WaitGroup
	// 这里为了测试下rpc server限流的拦截器是否有用
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rpcResp, err = l.svcCtx.Bing.Work(l.ctx, &bingclient.Request{})
			if err != nil {
				l.Info(err.Error())
				return
			}
		}()
	}
	wg.Wait()
	return &types.Response{
		Msg: rpcResp.Pong,
	}, nil
}
