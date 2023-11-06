package logic

import (
	"Thinkphoto/server/follow/internal/svc"
	"Thinkphoto/server/follow/pb/pb"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注
func (l *FollowLogic) Follow(in *pb.FollowRequest) (*pb.FollowResponse, error) {
	followerId := in.GetFollowerId()   // 关注者
	followingId := in.GetFollowingId() // 被关注者

	var err error
	// 在业务逻辑中调用事务操作
	err = l.svcCtx.Conn.Transact(func(session sqlx.Session) error {
		err = l.svcCtx.TFollowModel.FollowWithSession(session, followerId, followingId, 1)
		if err != nil {
			return err
		}
		err = l.svcCtx.TFollowCountModel.FollowCountWithSession(session, followerId, followingId, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logx.Errorf("follow fail:%v", err)
		return nil, err
	}
	// 关注成功
	return &pb.FollowResponse{}, nil
}
