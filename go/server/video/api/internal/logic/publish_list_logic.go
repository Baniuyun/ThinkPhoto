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

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListReq) (*types.PublishListResp, error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	if userId == 0 {
		return nil, xcode.NoLogin
	}
	list, err := l.svcCtx.VideoServer.GetPublishVideoList(l.ctx, &video.GetPublishVideoListRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	n := len(list.VideoList)
	resp := &types.PublishListResp{}
	for i := 0; i < n; i++ {
		resp.VideoList = append(resp.VideoList, types.VideoSimple{
			VideoId:       list.VideoList[i].VideoId,
			Information:   list.VideoList[i].Information,
			CoverURL:      list.VideoList[i].CoverUrl,
			FavoriteCount: list.VideoList[i].FavoriteCount})
	}
	return resp, nil
}
