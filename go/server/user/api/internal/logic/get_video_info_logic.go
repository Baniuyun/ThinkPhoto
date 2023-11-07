package logic

import (
	"Thinkphoto/server/follow/follow"
	"Thinkphoto/server/user/rpc/user"
	"context"

	"Thinkphoto/server/user/api/internal/svc"
	"Thinkphoto/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoInfoLogic {
	return &GetVideoInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoInfoLogic) GetVideoInfo(req *types.InfoReq) (*types.InfoResp, error) {
	// todo: add your logic here and delete this line
	id, err := l.svcCtx.UserServer.FindById(l.ctx, &user.FindByIdRequest{
		UserId: req.UserId,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.UserServer.FindById err: %v", err)
		return nil, err
	}
	followerCount, err := l.svcCtx.FollowerServer.FollowCount(l.ctx, &follow.FollowCountRequest{
		FollowerId: req.UserId,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.FollowerServer.FollowCount err: %v", err)
	}
	followingCount, err := l.svcCtx.FollowerServer.FansCount(l.ctx, &follow.FansCountRequest{
		FollowingId: req.UserId,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.FollowerServer.FollowCount err: %v", err)
	}
	follow, err := l.svcCtx.FollowerServer.IsFollow(l.ctx, &follow.IsFollowRequest{
		FollowerId:  req.UserId,
		FollowingId: req.UserId,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.FollowerServer.IsFollow err: %v", err)
	}
	resp := &types.InfoResp{
		Avatar:         id.Avatar,
		FollowerCount:  followerCount.Count,
		FollowingCount: followingCount.Count,
		IsFollow:       follow.Res,
		UserID:         id.UserId,
		UserName:       id.Username,
	}
	return resp, nil

}
