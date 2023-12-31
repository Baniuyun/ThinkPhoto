// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"Thinkphoto/server/video/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/video/list",
				Handler: getVideoListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/video/publishvideo",
				Handler: publishVideoHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: getVideoInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search",
				Handler: searchVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/publishtoken",
				Handler: PublishTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/publishlist",
				Handler: publishListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/publishvideo",
				Handler: deleteVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/favoritelist",
				Handler: favoriteListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/favorite",
				Handler: favoriteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/video"),
	)
}
