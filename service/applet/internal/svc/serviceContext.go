package svc

import (
	"BuzzBox/pkg/interceptors"
	"BuzzBox/service/applet/internal/config"
	"BuzzBox/service/user/rpc/userClient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
	UserRPC  userClient.User // 用户服务
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 自定义拦截器
	// 示例：添加一个自定义的拦截器
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:   c,
		BizRedis: redis.MustNewRedis(c.BizRedis),
		UserRPC:  userClient.NewUser(userRPC),
	}
}
