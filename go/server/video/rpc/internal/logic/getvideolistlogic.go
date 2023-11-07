package logic

import (
	"Thinkphoto/server/video/rpc/internal/model"
	"Thinkphoto/server/video/rpc/internal/svc"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取视频列表
func (l *GetVideoListLogic) GetVideoList(in *video.GetVideoListRequest) (*video.GetVideoListResponse, error) {
	// todo: add your logic here and delete this line
	tag := in.Tag
	var tVideoList []*model.TVideo
	var err error

	if tag == 0 {
		// 表示获取全部类型视频
		tVideoList, err = l.svcCtx.TVideoModel.FindListAll(context.Background())
	} else {
		// 获取当前 tag 分类视频
		tVideoList, err = l.svcCtx.TVideoModel.FindListByTag(context.Background(), tag)
	}
	if err != nil {
		logx.Errorf("tVideo Get err:%v", err)
		return nil, err
	}
	var videoList []*video.VideoInfo
	for i := 0; i < len(tVideoList); i++ {
		videoList = append(videoList, ConvertToVideoInfo(tVideoList[i], false))
	}
	return &video.GetVideoListResponse{VideoList: videoList}, nil
}
