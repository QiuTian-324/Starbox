package logic

import (
	"context"

	"BuzzBox/service/user/rpc/internal/svc"
	"BuzzBox/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *user.SendSmsRequest) (*user.SendSmsResponse, error) {
	// todo: add your logic here and delete this line

	return &user.SendSmsResponse{}, nil
}
