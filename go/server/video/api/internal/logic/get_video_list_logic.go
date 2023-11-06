package logic

import (
	"Thinkphoto/server/video/rpc/pb/video"
	"context"

	"Thinkphoto/server/video/api/internal/svc"
	"Thinkphoto/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoListLogic) GetVideoList(req *types.VideoListReq) (*types.VideoListResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.VideoServer.GetVideoList(l.ctx, &video.GetVideoListRequest{
		Tag:    req.Tag,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	n := len(list.VideoList)
	resp := &types.VideoListResp{}
	for i := 0; i < n; i++ {
		resp.VideoList = append(resp.VideoList, types.Video{
			User: types.User{
				UserId:   list.VideoList[i].AuthorId,
				IsFollow: list.VideoList[i].IsFavorite,
				UserName: list.VideoList[i].AuthorName,
				Avator:   list.VideoList[i].Avatar,
			},
			CommentCount:  list.VideoList[i].CommentCount,
			CoverURL:      list.VideoList[i].CoverUrl,
			FavoriteCount: list.VideoList[i].FavoriteCount,
			VideoId:       list.VideoList[i].VideoId,
			IsFavorite:    list.VideoList[i].IsFavorite,
			PlayURL:       list.VideoList[i].PlayUrl,
			Information:   list.VideoList[i].Information,
		})
	}
	return resp, nil
}
