package main

import (
	"Thinkphoto/pkg/cors"
	"Thinkphoto/pkg/xcode"
	"Thinkphoto/server/webapp/internal/config"
	"Thinkphoto/server/webapp/internal/handler"
	"Thinkphoto/server/webapp/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "server/webapp/etc/webapp.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf) //, rest.WithNotAllowedHandler(cors.NewCorsMiddleware().Handler()))
	server.Use(cors.NewCorsMiddleware().Handle)
	//server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误处理方法
	httpx.SetErrorHandler(xcode.ErrHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
