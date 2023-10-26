package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/pb"
)

type GetPublishVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishVideoListLogic {
	return &GetPublishVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户上传视频列表
func (l *GetPublishVideoListLogic) GetPublishVideoList(in *pb.GetPublishVideoListRequest) (*pb.GetPublishVideoListResponse, error) {
	// todo: add your logic here and delete this line
	return &pb.GetPublishVideoListResponse{}, nil
}
