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
	userId := in.GetFollowerId()

	tFollowCount, err := l.svcCtx.TFollowCountModel.FindOneByUserId(context.Background(), userId)
	if err != nil {
		logx.Errorf("FollowCount get fail:%v", err)
		return nil, err
	}

	count := tFollowCount.FollowingCount
	return &pb.CountResponse{Count: count}, nil
}
