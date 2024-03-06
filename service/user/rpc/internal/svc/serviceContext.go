package svc

import (
	"BuzzBox/service/user/rpc/internal/config"
	"BuzzBox/service/user/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
	}
}
