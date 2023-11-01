package logic

import (
	"context"

	"Thinkphoto/server/follow/internal/svc"
	"Thinkphoto/server/follow/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FansCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFansCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FansCountLogic {
	return &FansCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 粉丝数
func (l *FansCountLogic) FansCount(in *pb.FansCountRequest) (*pb.CountResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CountResponse{}, nil
}
