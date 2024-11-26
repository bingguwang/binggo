package logic

import (
	"context"

	"binggo/zero/zero-study/dtmDemo/product/rpc/internal/svc"
	"binggo/zero/zero-study/dtmDemo/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	// todo: add your logic here and delete this line

	return &product.DetailResponse{}, nil
}
