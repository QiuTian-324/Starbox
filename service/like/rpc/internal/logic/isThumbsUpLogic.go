package logic

import (
	"context"

	"BuzzBox/service/like/rpc/internal/svc"
	"BuzzBox/service/like/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsThumbsUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsThumbsUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsThumbsUpLogic {
	return &IsThumbsUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsThumbsUpLogic) IsThumbsUp(in *pb.IsThumbsUpRequest) (*pb.IsThumbsUpResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.IsThumbsUpResponse{}, nil
}
