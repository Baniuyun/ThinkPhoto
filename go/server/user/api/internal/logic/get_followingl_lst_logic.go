package logic

import (
	"Thinkphoto/server/follow/pb/pb"
	"Thinkphoto/server/user/rpc/pb/service"
	"context"

	"Thinkphoto/server/user/api/internal/svc"
	"Thinkphoto/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowinglLstLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowinglLstLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowinglLstLogic {
	return &GetFollowinglLstLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowinglLstLogic) GetFollowinglLst(req *types.FollowinglistReq) (*types.FollowinglistResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.FollowerServer.FansList(l.ctx, &pb.FansListRequest{
		FollowingId: req.UserID,
		Cursor:      0,
		PageSize:    100,
	})
	if err != nil {
		return nil, err
	}
	resp := &types.FollowinglistResp{}
	n := len(list.Items)
	for i := 0; i < n; i++ {
		id, err := l.svcCtx.UserServer.FindById(l.ctx, &service.FindByIdRequest{UserId: req.UserID})
		if err != nil {
			logx.Errorf("l.svcCtx.UserServer.FindById", err)
			return nil, err
		}
		resp.UserList = append(resp.UserList, types.UserSimple{
			Avatar: id.Avatar,
			//IsFollow: list.Items[i].FollowerId,
			UserID:   list.Items[i].FollowingId,
			UserName: id.Username,
		})
	}
	return resp, nil
}
