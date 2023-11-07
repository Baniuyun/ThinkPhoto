package svc

import (
	"Thinkphoto/pkg/interceptor"
	"Thinkphoto/server/follow/follow"
	"Thinkphoto/server/user/api/internal/config"
	"Thinkphoto/server/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	FollowerServer follow.Follow
	UserServer     user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	FollowRPC := zrpc.MustNewClient(c.FollowRPC, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorInterceptor()))
	UserRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:         c,
		FollowerServer: follow.NewFollow(FollowRPC),
		UserServer:     user.NewUser(UserRPC),
	}
}
