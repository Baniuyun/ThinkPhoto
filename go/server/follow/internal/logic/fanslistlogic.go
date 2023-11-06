package logic

import (
	"Thinkphoto/server/follow/follow"
	"Thinkphoto/server/follow/internal/model"
	"context"

	"Thinkphoto/server/follow/internal/svc"
	"Thinkphoto/server/follow/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FansListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FansListLogic {
	return &FansListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 粉丝列表
func (l *FansListLogic) FansList(in *pb.FansListRequest) (*pb.FansListResponse, error) {
	followingId := in.GetFollowingId()
	cursor := in.GetCursor()
	pageSize := in.GetPageSize()

	fansList, err := l.svcCtx.TFollowModel.FindFansList(context.Background(), followingId, cursor, pageSize)
	if err != nil {
		logx.Errorf("FindFansList fail:%v", err)
		return nil, err
	}

	var fansItemList []*follow.FansItem

	for i := 0; i < len(fansList); i++ {
		tFollow := ConvertToFansItem(fansList[i])
		fansItemList = append(fansItemList, tFollow)
	}

	return &pb.FansListResponse{Items: fansItemList, Cursor: cursor + 1}, nil
}

func ConvertToFansItem(tFollow *model.TFollow) *follow.FansItem {
	return &follow.FansItem{
		FollowingId: tFollow.FollowingId,
		FollowerId:  tFollow.FollowerId,
		CreateTime:  tFollow.CreateTime.Unix(),
	}
}
