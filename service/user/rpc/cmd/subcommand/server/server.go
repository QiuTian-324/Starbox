package server

import (
	"BuzzBox/service/user/rpc/internal/config"
	"BuzzBox/service/user/rpc/internal/server"
	"BuzzBox/service/user/rpc/internal/svc"
	"BuzzBox/service/user/rpc/user"
	"context"
	"flag"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "../etc/user.yaml", "the config file")

// StartServe 启动用户 RPC 服务器.
func StartServe(cmd *cobra.Command, args []string) {
	flag.Parse()

	var c config.Config
	// 从配置文件加载配置信息
	conf.MustLoad(*configFile, &c)
	// 创建服务上下文
	ctx := svc.NewServiceContext(c)

	// 创建 RPC 服务器
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// 注册用户服务
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		// 在开发模式或测试模式下注册服务反射
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 启动 RPC 服务器
	logc.Infof(context.Background(), "启动rpc服务器: %s...\n", c.ListenOn)
	s.Start()
}
