package svc

import (
	"BuzzBox/pkg/interceptors"
	"BuzzBox/service/article/api/internal/config"
	"BuzzBox/service/article/rpc/article"
	"BuzzBox/service/user/rpc/userClient"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/zrpc"
)

const (
	// 默认连接超时时间
	defaultOssConnectTimeout = 1
	// 默认读写超时时间
	defaultOssReadWriteTimeout = 3
)

type ServiceContext struct {
	Config     config.Config
	OssClient  *oss.Client
	ArticleRPC article.Article
	UserRPC    userClient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	if c.Oss.ConnectTimeout == 0 {
		c.Oss.ConnectTimeout = defaultOssConnectTimeout
	}
	if c.Oss.ReadWriteTimeout == 0 {
		c.Oss.ReadWriteTimeout = defaultOssReadWriteTimeout
	}

	c.RestConf.MaxBytes = c.MaxBytes

	oc, err := oss.New(c.Oss.Endpoint, c.Oss.AccessKeyId, c.Oss.AccessKeySecret,
		oss.Timeout(c.Oss.ConnectTimeout, c.Oss.ReadWriteTimeout))
	if err != nil {
		panic(err)
	}

	//  创建RPC客户端，客户端拦截器
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	articleRPC := zrpc.MustNewClient(c.ArticleRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:     c,
		OssClient:  oc,
		ArticleRPC: article.NewArticle(articleRPC),
		UserRPC:    userClient.NewUser(userRPC),
	}
}
