package main

import (
	"Thinkphoto/pkg/interceptor"
	"Thinkphoto/server/user/rpc/internal/config"
	"Thinkphoto/server/user/rpc/internal/server"
	"Thinkphoto/server/user/rpc/internal/svc"
	svr "Thinkphoto/server/user/rpc/pb/service"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/zeromicro/go-zero/core/conf"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		svr.RegisterUserServer(grpcServer, server.NewUserServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 自定义拦截器
	s.AddUnaryInterceptors(interceptor.ServerErrorInterceptor())
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
