package logic

import (
	"context"

	"Thinkphoto/server/follow/internal/svc"
	"Thinkphoto/server/follow/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowLogic {
	return &IsFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 是否关注
func (l *IsFollowLogic) IsFollow(in *pb.IsFollowRequest) (*pb.IsFollowResponse, error) {
	followerId := in.GetFollowerId()   // 关注者id
	followingId := in.GetFollowingId() // 被关注者

	follow, err := l.svcCtx.TFollowModel.FindOneByFollowerIdFollowingId(context.Background(), followerId, followingId)

	if err != nil {
		logx.Errorf("find by follower_id and following_id fail:%v", err)
		return nil, err
	}

	isFollow := follow.IsFollow

	return &pb.IsFollowResponse{Res: isFollow == 1}, nil
}
