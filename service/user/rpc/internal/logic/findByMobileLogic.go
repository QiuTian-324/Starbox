package logic

import (
	"BuzzBox/pkg/encrypt"
	"BuzzBox/service/user/rpc/pkg/user_rpc_code"
	"context"

	"BuzzBox/service/user/rpc/internal/svc"
	"BuzzBox/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FindByMobile 查询手机号
func (l *FindByMobileLogic) FindByMobile(in *user.FindByMobileRequest) (*user.FindByMobileResponse, error) {
	u, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		return nil, user_rpc_code.ErrMobileNotExist
	}
	if u.Mobile != in.Mobile {
		return nil, user_rpc_code.ErrNotRegister
	}

	return &user.FindByMobileResponse{
		UserId:   u.Id,
		Username: u.Username,
		Avatar:   u.Avatar,
		Mobile:   u.Mobile,
		CreateAt: encrypt.FormatCreateTime(u.CreateTime),
		UpdateAt: encrypt.FormatUpdateTime(u.UpdateTime),
	}, nil
}
