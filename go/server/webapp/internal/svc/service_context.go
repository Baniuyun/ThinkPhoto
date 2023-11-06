package svc

import (
	"Thinkphoto/pkg/interceptor"
	"Thinkphoto/server/user/rpc/user"
	"Thinkphoto/server/webapp/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorInterceptor()))
	return &ServiceContext{
		Config: c,
		User:   user.NewUser(userRPC),
	}
}
