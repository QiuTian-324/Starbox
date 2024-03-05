package logic

import (
	"BuzzBox/service/user/rpc/internal/svc"
	"BuzzBox/service/user/rpc/pkg/user_code"
	"BuzzBox/service/user/rpc/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 查找出用户信息，然后对比密码
	userInfo, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		return nil, err
	}
	if userInfo.Password != in.Password {
		return nil, user_code.ErrPassword
	}
	return &user.LoginResponse{
		Id: userInfo.Id,
	}, nil
}
