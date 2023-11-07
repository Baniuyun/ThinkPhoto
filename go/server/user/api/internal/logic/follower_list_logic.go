package logic

import (
	"Thinkphoto/server/follow/follow"
	"Thinkphoto/server/user/rpc/pb/service"
	"context"

	"Thinkphoto/server/user/api/internal/svc"
	"Thinkphoto/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerlistReq) (*types.FollowerlistResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.FollowerServer.FollowList(l.ctx, &follow.FollowListRequest{
		Id:         req.UserID,
		FollowerId: req.UserID,
		Cursor:     0,
		PageSize:   100,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.FollowerServer.FollowList", err)
		return nil, err
	}
	n := len(list.Items)
	resp := &types.FollowerlistResp{}
	for i := 0; i < n; i++ {
		id, err := l.svcCtx.UserServer.FindById(l.ctx, &service.FindByIdRequest{UserId: req.UserID})
		if err != nil {
			logx.Errorf("l.svcCtx.UserServer.FindById", err)
			return nil, err
		}
		resp.UserList = append(resp.UserList, types.UserSimple{
			Avatar: id.Avatar,
			//IsFollow: list.Items[i].FollowerId,
			UserID:   list.Items[i].FollowerId,
			UserName: id.Username,
		})
	}
	return resp, nil
}
