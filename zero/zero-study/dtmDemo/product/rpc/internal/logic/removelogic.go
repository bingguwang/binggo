package logic

import (
	"context"

	"binggo/zero/zero-study/dtmDemo/product/rpc/internal/svc"
	"binggo/zero/zero-study/dtmDemo/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *product.RemoveRequest) (*product.RemoveResponse, error) {
	// todo: add your logic here and delete this line

	return &product.RemoveResponse{}, nil
}
