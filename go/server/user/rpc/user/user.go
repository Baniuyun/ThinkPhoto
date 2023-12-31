// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"Thinkphoto/server/user/rpc/pb/service"
	"context"


	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FindByIdRequest        = service.FindByIdRequest
	FindByIdResponse       = service.FindByIdResponse
	FindByUserNameRequest  = service.FindByUserNameRequest
	FindByUserNameResponse = service.FindByUserNameResponse
	RegisterRequest        = service.RegisterRequest
	RegisterResponse       = service.RegisterResponse

	User interface {
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error)
		FindByUserName(ctx context.Context, in *FindByUserNameRequest, opts ...grpc.CallOption) (*FindByUserNameResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*FindByIdResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.FindById(ctx, in, opts...)
}

func (m *defaultUser) FindByUserName(ctx context.Context, in *FindByUserNameRequest, opts ...grpc.CallOption) (*FindByUserNameResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.FindByUserName(ctx, in, opts...)
}
