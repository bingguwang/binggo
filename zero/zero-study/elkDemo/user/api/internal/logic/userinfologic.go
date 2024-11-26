package logic

import (
	"context"
	"github.com/google/martian/log"
	"github.com/zeromicro/go-zero/core/logx"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/common/xerr"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	logx.Info("哈哈哈来了")
	log.Infof("logDemo.Infof哈哈哈来了")
	log.Errorf("logDemo.Errorf哈哈哈来了error")
	logx.Error("哈哈哈来了error")

	if err == nil {
		return nil, xerr.NewErrCodeMsg(500, "用户查询失败")
	}

	return &types.UserInfoResponse{}, nil
}
