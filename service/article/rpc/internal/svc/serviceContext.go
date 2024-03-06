package svc

import (
	"BuzzBox/service/article/rpc/internal/config"
	"BuzzBox/service/article/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.ArticleModel // 添加ArticleModel
	BizRedis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})

	if err != nil {
		logx.Error("redis连接失败")
		panic(err)
	}

	return &ServiceContext{
		Config:       c,
		BizRedis:     rds,
		ArticleModel: model.NewArticleModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
	}
}
