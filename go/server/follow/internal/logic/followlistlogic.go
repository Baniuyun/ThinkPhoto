package logic

import (
	"Thinkphoto/server/follow/follow"
	"context"

	"Thinkphoto/server/follow/internal/svc"
	"Thinkphoto/server/follow/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注列表
func (l *FollowListLogic) FollowList(in *pb.FollowListRequest) (*pb.FollowListResponse, error) {
	followerId := in.GetFollowerId()
	cursor := in.GetCursor()
	pageSize := in.GetPageSize()

	tFollowList, err := l.svcCtx.TFollowModel.FindFollowList(context.Background(), followerId, cursor, pageSize)
	if err != nil {
		logx.Errorf("FindFollowList fail:%v", err)
		return nil, err
	}

	var followList []*follow.FollowItem

	for i := 0; i < len(tFollowList); i++ {
		followItem := &follow.FollowItem{
			Id:         tFollowList[i].Id,
			FollowerId: tFollowList[i].FollowerId,
			FansCount:  tFollowList[i].FansCount,
			CreateTime: tFollowList[i].CreateTime,
		}
		followList = append(followList, followItem)
	}

	return &pb.FollowListResponse{Items: followList, Cursor: cursor + 1, IsEnd: len(followList) >= int(pageSize), Id: followerId}, nil
}
