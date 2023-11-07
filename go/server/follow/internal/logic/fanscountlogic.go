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
	userId := in.GetFollowingId() // 用户id

	tFollowCount, err := l.svcCtx.TFollowCountModel.FindOneByUserId(context.Background(), userId)

	if err != nil {
		logx.Errorf("TFollowCount Get err:%v", err)
		return nil, err
	}

	count := tFollowCount.FollowerCount // 粉丝数

	return &pb.CountResponse{Count: count}, nil
}
