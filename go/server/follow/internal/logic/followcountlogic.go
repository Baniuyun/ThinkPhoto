package logic

import (
	"context"

	"Thinkphoto/server/follow/internal/svc"
	"Thinkphoto/server/follow/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowCountLogic {
	return &FollowCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注数
func (l *FollowCountLogic) FollowCount(in *pb.FollowCountRequest) (*pb.CountResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CountResponse{}, nil
}
