// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"BuzzBox/service/user/rpc/internal/logic"
	"BuzzBox/service/user/rpc/internal/svc"
	"BuzzBox/service/user/rpc/user"
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

func (s *UserServer) FindById(ctx context.Context, in *user.FindByIdRequest) (*user.FindByIdResponse, error) {
	l := logic.NewFindByIdLogic(ctx, s.svcCtx)
	return l.FindById(in)
}

func (s *UserServer) FindByMobile(ctx context.Context, in *user.FindByMobileRequest) (*user.FindByMobileResponse, error) {
	l := logic.NewFindByMobileLogic(ctx, s.svcCtx)
	return l.FindByMobile(in)
}

func (s *UserServer) GetUserInfoByID(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewGetUserInfoByIDLogic(ctx, s.svcCtx)
	return l.GetUserInfoByID(in)
}

func (s *UserServer) SendSms(ctx context.Context, in *user.SendSmsRequest) (*user.SendSmsResponse, error) {
	l := logic.NewSendSmsLogic(ctx, s.svcCtx)
	return l.SendSms(in)
}
