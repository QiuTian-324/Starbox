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

func StartServe(cmd *cobra.Command, args []string) {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logc.Infof(context.Background(), "启动rpc服务器: %s...\n", c.ListenOn)
	s.Start()
}
