package logic

import (
	"BuzzBox/service/article/api/internal/pkg/article_api_code"
	"BuzzBox/service/article/rpc/pb"
	"context"
	"encoding/json"

	"BuzzBox/service/article/api/internal/svc"
	"BuzzBox/service/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishArticleLogic {
	return &PublishArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var (
	minContent = 2
)

func (l *PublishArticleLogic) PublishArticle(req *types.PublishRequest) (resp *types.PublishResponse, err error) {
	// todo: add your logic here and delete this line
	if len(req.Title) == 0 {
		return nil, article_api_code.ArticleTitleEmpty
	}
	if len(req.Content) < minContent {
		return nil, article_api_code.ArticleContentTooFewWords
	}
	if len(req.Cover) == 0 {
		return nil, article_api_code.ArticleCoverEmpty
	}

	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		logx.Errorf("获取token失败: %v", err)
		return nil, article_api_code.NotLoginError
	}

	if userId == 0 {
		return nil, article_api_code.NotLoginError
	}

	// 发布文章
	article, err := l.svcCtx.ArticleRPC.PublishArticle(l.ctx, &pb.PublishArticleRequest{
		UserId:      userId,
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		Cover:       req.Cover,
	})
	if err != nil {
		return nil, article_api_code.ArticlePublishFailed
	}

	return &types.PublishResponse{ArticleId: article.ArticleId}, nil
}
