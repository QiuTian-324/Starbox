package logic

import (
	"BuzzBox/pkg/encrypt"
	"context"

	"BuzzBox/service/user/rpc/internal/svc"
	"BuzzBox/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIDLogic {
	return &GetUserInfoByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByIDLogic) GetUserInfoByID(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	mobile, _ := encrypt.DecMobile(u.Mobile)

	return &user.UserInfoResponse{
		UserId:   u.Id,
		Username: u.Username,
		Avatar:   u.Avatar,
		Mobile:   mobile,
		CreateAt: encrypt.FormatCreateTime(u.CreateTime),
		UpdateAt: encrypt.FormatUpdateTime(u.UpdateTime),
	}, nil
}
