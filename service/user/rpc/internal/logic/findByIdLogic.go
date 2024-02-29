package logic

import (
	"context"

	"BuzzBox/service/user/rpc/internal/svc"
	"BuzzBox/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *user.FindByIdRequest) (*user.FindByIdResponse, error) {
	// todo: add your logic here and delete this line

	return &user.FindByIdResponse{}, nil
}
