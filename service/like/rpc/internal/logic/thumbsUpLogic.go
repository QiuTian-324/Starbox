package logic

import (
	"context"

	"BuzzBox/service/like/rpc/internal/svc"
	"BuzzBox/service/like/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbsUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbsUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbsUpLogic {
	return &ThumbsUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbsUpLogic) ThumbsUp(in *pb.ThumbsUpRequest) (*pb.ThumbsUpResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ThumbsUpResponse{}, nil
}
