package logic

import (
	"context"
	"fmt"
	"rpc/internal/model"

	"rpc/internal/svc"
	"rpc/pb"

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
func (l *GetVideoListLogic) GetVideoList(in *pb.GetVideoListRequest) (*pb.GetVideoListResponse, error) {
	//currentTime := in.CurrentTime
	//tokenUserId := in.TokenUserId
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
		fmt.Println("Get t_video fail: ", err)
		return nil, err
	}

	var videoList []*pb.VideoInfo
	for i := 0; i < len(tVideoList); i++ {
		videoInfo := ConvertToVideoInfo(tVideoList[i])
		videoList[i] = videoInfo
	}

	return &pb.GetVideoListResponse{VideoList: videoList}, nil
}
