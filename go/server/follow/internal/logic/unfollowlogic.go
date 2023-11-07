package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"Thinkphoto/server/follow/internal/svc"
	"Thinkphoto/server/follow/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFollowLogic {
	return &UnFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消关注
func (l *UnFollowLogic) UnFollow(in *pb.UnFollowRequest) (*pb.UnFollowResponse, error) {
	followerId := in.GetFollowerId()   // 取关者
	followingId := in.GetFollowingId() // 被取关者

	var err error
	// 在业务逻辑中调用事务操作
	err = l.svcCtx.Conn.Transact(func(session sqlx.Session) error {
		err = l.svcCtx.TFollowModel.FollowWithSession(session, followerId, followingId, 0)
		if err != nil {
			return err
		}
		err = l.svcCtx.TFollowCountModel.FollowCountWithSession(session, followerId, followingId, 0)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logx.Errorf("follow fail:%v", err)
		return nil, err
	}
	// 取关成功
	return &pb.UnFollowResponse{}, nil
}
