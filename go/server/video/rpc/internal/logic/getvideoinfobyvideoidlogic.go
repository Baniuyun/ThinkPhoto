package logic

import (
	"Thinkphoto/server/video/rpc/code"
	"Thinkphoto/server/video/rpc/internal/model"
	"Thinkphoto/server/video/rpc/internal/svc"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoInfoBYVideoIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoInfoBYVideoIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoInfoBYVideoIdLogic {
	return &GetVideoInfoBYVideoIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据视频id获取视频信息
func (l *GetVideoInfoBYVideoIdLogic) GetVideoInfoBYVideoId(in *video.GetVideoInfoReq) (*video.GetVideoInfoResp, error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.TVideoModel.FindOne(l.ctx, in.VideoId)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, code.VideoNil
		}
		logx.Errorf("l.svcCtx.TVideoModel.FindOne err: %v", err)
		return nil, err
	}
	_, err = l.svcCtx.TVideoLikeRecordModel.FindLikeRelation(l.ctx, one.Id, in.UserId)
	IsFavotite := true
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			IsFavotite = false
		} else {
			logx.Errorf("l.svcCtx.TVideoLikeRecordModel.FindLikeRelation err: %v", err)
			return nil, err
		}
	}
	return &video.GetVideoInfoResp{Info: ConvertToVideoInfo(one, IsFavotite)}, nil
}

func ConvertToVideoInfo(tVideo *model.TVideo, IsFavorite bool) *video.VideoInfo {

	return &video.VideoInfo{
		VideoId:       tVideo.Id,
		AuthorId:      tVideo.AuthorId,
		PlayUrl:       tVideo.PlayUrl,
		CoverUrl:      "",
		FavoriteCount: tVideo.FavoriteCount,
		CommentCount:  0, //tVideo.CommentCount,
		IsFavorite:    IsFavorite,
		AuthorName:    tVideo.AuthorName,
		Information:   tVideo.Information,
		Tag:           int32(tVideo.Tag),
		Time:          tVideo.Time,
	}
}
