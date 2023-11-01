package main

import (
	"Thinkphoto/pkg/interceptor"
	"Thinkphoto/server/Zinc/internal/config"
	"Thinkphoto/server/Zinc/internal/server"
	"Thinkphoto/server/Zinc/internal/svc"
	"Thinkphoto/server/Zinc/pb/zinc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "server/Zinc/etc/zinc.yaml", "the config file")

func main() {
	flag.Parse()
	//fmt.Println(os.Getwd())
	var c config.Config
	conf.MustLoad(*configFile, &c)
	//logx.DisableStat()
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		zinc.RegisterZincSearchServer(grpcServer, server.NewZincSearchServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(interceptor.ServerErrorInterceptor())
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
