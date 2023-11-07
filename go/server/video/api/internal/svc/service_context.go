package svc

import (
	"Thinkphoto/pkg/interceptor"
	"Thinkphoto/server/Zinc/zincsearch"
	"Thinkphoto/server/video/api/internal/config"
	"Thinkphoto/server/video/rpc/videoservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	VideoServer  videoservice.VideoService
	SearchServer zincsearch.ZincSearch
}

func NewServiceContext(c config.Config) *ServiceContext {
	videoRPC := zrpc.MustNewClient(c.VideoRPC, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorInterceptor()))
	searchRPC := zrpc.MustNewClient(c.SearchRPC, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:       c,
		VideoServer:  videoservice.NewVideoService(videoRPC),
		SearchServer: zincsearch.NewZincSearch(searchRPC),
	}
}
