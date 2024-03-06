package logic

import (
	"context"

	"BuzzBox/service/article/rpc/internal/svc"
	"BuzzBox/service/article/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleListLogic) GetArticleList(in *pb.GetArticleListRequest) (*pb.GetArticleListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetArticleListResponse{}, nil
}
