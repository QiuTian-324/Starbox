package logic

import (
	"BuzzBox/service/user/model"
	"BuzzBox/service/user/rpc/internal/svc"
	"BuzzBox/service/user/rpc/pkg/user_code"
	"BuzzBox/service/user/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {

	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:   in.Username,
		Password:   in.Password,
		Mobile:     in.Mobile,
		Avatar:     in.Avatar,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, user_code.ErrRegisterFail
	}
	userId, err := ret.LastInsertId()
	if err != nil {
		logx.Errorf("用户ID插入失败: %v", err)
		return nil, user_code.ErrUserIDInsertFail
	}

	return &user.RegisterResponse{Id: userId}, nil
}
