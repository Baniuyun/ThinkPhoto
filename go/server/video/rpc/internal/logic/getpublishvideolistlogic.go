package logic

import (
	"Thinkphoto/server/video/rpc/code"
	"Thinkphoto/server/video/rpc/internal/model"
	"Thinkphoto/server/video/rpc/internal/svc"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishVideoListLogic {
	return &GetPublishVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户上传视频列表
func (l *GetPublishVideoListLogic) GetPublishVideoList(in *video.GetPublishVideoListRequest) (*video.GetPublishVideoListResponse, error) {
	// todo: add your logic here and delete this line
	if in.UserId == 0 {
		return nil, code.UserIdEmpty
	}
	// 获取上传视频列表
	tVideoList, err := l.svcCtx.TVideoModel.FindListByAuthorId(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("FIndListById err:%v", err)
		return nil, err
	}
	var videoList []*video.VideoSimple
	for i := 0; i < len(tVideoList); i++ {
		videoList = append(videoList, ConvertToVideoSimple(tVideoList[i]))
	}
	return &video.GetPublishVideoListResponse{VideoList: videoList}, nil
}

// 类型转换
func ConvertToVideoSimple(tvideo *model.TVideo) *video.VideoSimple {
	return &video.VideoSimple{
		VideoId:       tvideo.Id,
		CoverUrl:      "",
		FavoriteCount: tvideo.FavoriteCount,
		Information:   tvideo.Information,
		Time:          tvideo.Time,
	}
}
