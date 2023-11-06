package logic

import (
	"Thinkphoto/server/follow/follow"
	"context"

	"Thinkphoto/server/user/api/internal/svc"
	"Thinkphoto/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResp, err error) {
	// todo: add your logic here and delete this line
	//如果不能返回状态码
	//isFollow, err := l.svcCtx.FollowerServer.IsFollow(l.ctx, &follow.IsFollowRequest{
	//	FollowerId:  req.FollowerID,
	//	FollowingId: req.FollowerID,
	//})
	//if err != nil {
	//	logx.Errorf("l.svcCtx.FollowerServer.IsFollow err: %v",err)
	//	return nil, err
	//}
	//if isFollow.Res{
	//
	//}
	if req.Status == 0 {
		_, err := l.svcCtx.FollowerServer.Follow(l.ctx, &follow.FollowRequest{
			FollowerId:  req.FollowerID,
			FollowingId: req.FollowingID,
		})
		if err != nil {
			logx.Errorf("l.svcCtx.FollowerServer.Follow : %v", err)
			return nil, err
		}
	} else {
		_, err = l.svcCtx.FollowerServer.UnFollow(l.ctx, &follow.UnFollowRequest{
			FollowerId:  req.FollowerID,
			FollowingId: req.FollowingID,
		})
		if err != nil {
			logx.Errorf("l.svcCtx.FollowerServer.Follow : %v", err)
			return nil, err
		}
	}
	return
}
