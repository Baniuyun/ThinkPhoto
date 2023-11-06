package logic

import (
	"Thinkphoto/server/video/rpc/pb/video"
	"context"

	"Thinkphoto/server/video/api/internal/svc"
	"Thinkphoto/server/video/api/internal/types"

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
	id, err := l.svcCtx.VideoServer.GetVideoInfoBYVideoId(l.ctx, &video.GetVideoInfoReq{VideoId: req.VideoId})
	if err != nil {
		logx.Errorf("l.svcCtx.VideoServer.GetVideoInfoBYVideoId err%v", err)
		return nil, err
	}
	return &types.InfoResp{Video: types.Video{
		User: types.User{
			UserId:   id.Info.AuthorId,
			IsFollow: id.Info.IsFavorite,
			UserName: id.Info.AuthorName,
			Avator:   id.Info.Avatar,
		},
		CommentCount:  id.Info.CommentCount,
		CoverURL:      id.Info.CoverUrl,
		FavoriteCount: id.Info.FavoriteCount,
		VideoId:       id.Info.VideoId,
		IsFavorite:    id.Info.IsFavorite,
		PlayURL:       id.Info.PlayUrl,
		Information:   id.Info.Information,
	}}, nil
}
