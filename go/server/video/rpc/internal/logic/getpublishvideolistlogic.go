package logic

import (
	"Thinkphoto/server/video/rpc/code"
	"Thinkphoto/server/video/rpc/internal/model"
	"Thinkphoto/server/video/rpc/internal/svc"
	"Thinkphoto/server/video/rpc/pb"
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
func (l *GetPublishVideoListLogic) GetPublishVideoList(in *pb.GetPublishVideoListRequest) (*pb.GetPublishVideoListResponse, error) {
	userId := in.UserId
	//token := in.UserToken
	//
	//// 验证 token
	//if token == "" {
	//	return nil, nil
	//}
	if in.UserId == 0 {
		return nil, code.UserIdEmpty
	}
	// 获取上传视频列表
	tVideoList, err := l.svcCtx.TVideoModel.FindListById(context.Background(), userId)
	if err != nil {
		logx.Errorf("FIndListById err:%v", err)
		return nil, err
	}

	var videoList []*pb.VideoInfo
	for i := 0; i < len(tVideoList); i++ {
		videoInfo := ConvertToVideoInfo(tVideoList[i])
		videoList[i] = videoInfo
	}

	return &pb.GetPublishVideoListResponse{VideoList: videoList}, nil
}

// 类型转换
func ConvertToVideoInfo(tvideo *model.TVideo) *pb.VideoInfo {
	return &pb.VideoInfo{
		VideoId:       tvideo.Id,
		AuthorId:      tvideo.AuthorId,
		PlayUrl:       tvideo.PlayUrl,
		CoverUrl:      tvideo.CoverUrl,
		FavoriteCount: tvideo.FavoriteCount,
		CommentCount:  tvideo.CommentCount,
		IsFavorite:    int64ToBool(tvideo.IsFavorite),
		Title:         tvideo.Title.String,
		Information:   tvideo.Information,
		Tag:           int32(tvideo.Tag),
		Time:          tvideo.Time,
	}
}

func int64ToBool(value int64) bool {
	return value != 0
}
