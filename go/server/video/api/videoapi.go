package main

import (
	"Thinkphoto/pkg/cors"
	"Thinkphoto/pkg/xcode"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"Thinkphoto/server/video/api/internal/config"
	"Thinkphoto/server/video/api/internal/handler"
	"Thinkphoto/server/video/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "server/video/api/etc/videoapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	server.Use(cors.NewCorsMiddleware().Handle)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(xcode.ErrHandler)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
