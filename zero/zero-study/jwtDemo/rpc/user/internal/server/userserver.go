// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"binggo/zero/zero-study/jwtDemo/rpc/user/internal/logic"
	"binggo/zero/zero-study/jwtDemo/rpc/user/internal/svc"
	"binggo/zero/zero-study/jwtDemo/rpc/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}