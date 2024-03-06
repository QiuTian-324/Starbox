package logic

import (
	"BuzzBox/service/article/rpc/internal/model"
	"BuzzBox/service/article/rpc/internal/pkg/article_rpc_code"
	"BuzzBox/service/article/rpc/internal/pkg/types"
	"context"
	"time"

	"BuzzBox/service/article/rpc/internal/svc"
	"BuzzBox/service/article/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishArticleLogic {
	return &PublishArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishArticleLogic) PublishArticle(in *pb.PublishArticleRequest) (*pb.PublishArticleResponse, error) {
	// 如果已经登录，则将文章信息写入数据库
	insert, err := l.svcCtx.ArticleModel.Insert(l.ctx, &model.Article{
		Title:       in.Title,
		Content:     in.Content,
		Cover:       in.Cover,
		Description: in.Description,
		AuthorId:    in.UserId,
		Status:      types.ArticleStatusVisible,
		PublishTime: time.Now(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	})
	if err != nil {
		return nil, article_rpc_code.CodeArticlePublishFailed
	}
	id, err := insert.LastInsertId()
	if err != nil {
		return nil, article_rpc_code.CodeArticlePublishFailed
	}

	return &pb.PublishArticleResponse{ArticleId: id}, nil
}
