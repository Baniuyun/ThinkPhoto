package logic

import (
	"Thinkphoto/server/Zinc/pb/zinc"
	"Thinkphoto/server/video/api/internal/svc"
	"Thinkphoto/server/video/api/internal/types"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishVideoReq) (*types.PublishVideoResp, error) {
	// todo: add your logic here and delete this line
	publishVideo, err := l.svcCtx.VideoServer.PublishVideo(l.ctx, &video.PublishVideoRequest{
		Information: req.Information,
		Tag:         req.Tag,
		UserId:      req.UserId,
		Username:    req.Username,
		Avatar:      req.Avatar,
		Key:         req.Key,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.VideoServer.PublishVideo err %v", err)
		return nil, err
	}
	//fmt.Println("cjb")
	_, err = l.svcCtx.SearchServer.Doc(l.ctx, &zinc.Doc{
		VideoId:     publishVideo.Info.VideoId,
		Information: req.Information,
		UserName:    req.Username,
		UserId:      req.UserId,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.SearchServer.Doc err %v", err)
		return nil, err
	}
	//fmt.Println("cjb")
	return &types.PublishVideoResp{
		Video: types.Video{
			User: types.User{
				UserId:   publishVideo.Info.AuthorId,
				IsFollow: false,
				UserName: publishVideo.Info.AuthorName,
				Avator:   publishVideo.Info.Avatar,
			},
			CommentCount:  publishVideo.Info.CommentCount,
			CoverURL:      publishVideo.Info.CoverUrl,
			FavoriteCount: publishVideo.Info.FavoriteCount,
			VideoId:       publishVideo.Info.VideoId,
			IsFavorite:    publishVideo.Info.IsFavorite,
			PlayURL:       publishVideo.Info.PlayUrl,
			Information:   publishVideo.Info.Information,
		},
	}, nil
}
