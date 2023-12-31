package main

import (
	"Thinkphoto/pkg/cors"
	"flag"
	"fmt"

	"Thinkphoto/server/user/api/internal/config"
	"Thinkphoto/server/user/api/internal/handler"
	"Thinkphoto/server/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	server.Use(cors.NewCorsMiddleware().Handle)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
