package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublishTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishTokenLogic {
	return &GetPublishTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取上传回调凭证
func (l *GetPublishTokenLogic) GetPublishToken(in *pb.GetPublishTokenRequest) (*pb.GetPublishTokenResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPublishTokenResponse{}, nil
}
