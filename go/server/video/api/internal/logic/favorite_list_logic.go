package logic

import (
	"Thinkphoto/pkg/xcode"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"
	"encoding/json"

	"Thinkphoto/server/video/api/internal/svc"
	"Thinkphoto/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteVideoListReq) (*types.FavoriteVideoListResp, error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	if userId == 0 {
		return nil, xcode.NoLogin
	}
	list, err := l.svcCtx.VideoServer.GetFavoriteVideoList(l.ctx, &video.GetFavoriteVideoListRequest{UserId: req.UserId})
	if err != nil {
		logx.Errorf("l.svcCtx.VideoServer.GetFavoriteVideoList err %v", err)
		return nil, err
	}
	//resp.VideoList = list.VideoList
	resp := &types.FavoriteVideoListResp{}
	n := len(list.VideoList)
	for i := 0; i < n; i++ {
		resp.VideoList = append(resp.VideoList, types.VideoSimple{
			VideoId:       list.VideoList[i].VideoId,
			Information:   list.VideoList[i].Information,
			CoverURL:      list.VideoList[i].CoverUrl,
			FavoriteCount: list.VideoList[i].FavoriteCount})
	}
	return resp, nil
}
