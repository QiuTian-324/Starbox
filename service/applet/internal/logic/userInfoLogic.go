package logic

import (
	"BuzzBox/service/user/rpc/user"
	"context"
	"encoding/json"

	"BuzzBox/service/applet/internal/svc"
	"BuzzBox/service/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// 获取用户ID
	userID, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	//通过ID查找用户信息
	userInfo, err := l.svcCtx.UserRPC.GetUserInfoByID(l.ctx, &user.UserInfoRequest{
		Id: userID,
	})

	return &types.UserInfoResponse{
		ID:        userID,
		Username:  userInfo.Username,
		Avatar:    userInfo.Avatar,
		Mobile:    userInfo.Mobile,
		CreatedAt: userInfo.CreateAt,
		UpdatedAt: userInfo.UpdateAt,
	}, nil
}
