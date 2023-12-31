// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"Thinkphoto/server/webapp/internal/svc"
	"net/http"



	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/tpvideo"),
	)
}
