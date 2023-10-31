package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"rpc/internal/model"
	"rpc/internal/svc"
	"rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传视频
func (l *PublishVideoLogic) PublishVideo(in *pb.PublishVideoRequest) (*pb.PublishVideoResponse, error) {
	// todo: add your logic here and delete this line
	// 解析请求体
	var videoInfo pb.VideoInfo
	if err := json.Unmarshal([]byte(in.CallbackBody), &videoInfo); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return nil, err
	}

	// 生成视频地址和封面地址
	qiniu := l.svcCtx.Config.Qiniu
	videoInfo.PlayUrl = qiniu.Addr + qiniu.VideoPath + videoInfo.Title + ".mp4"
	videoInfo.CoverUrl = qiniu.Addr + qiniu.VideoPath + videoInfo.Title + ".jpg"

	tVideo := model.TVideo{
		AuthorId: videoInfo.AuthorId,
		PlayUrl:  videoInfo.PlayUrl,
		CoverUrl: videoInfo.CoverUrl,
		Title: sql.NullString{
			String: videoInfo.Title,
			Valid:  true,
		},
		Information: videoInfo.Information,
		Tag:         int64(videoInfo.Tag),
		Time:        videoInfo.Time,
	}

	res, err := l.svcCtx.TVideoModel.Insert(context.Background(), &tVideo)
	if err != nil {
		fmt.Println("insert fail: ", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	videoInfo.VideoId = id

	return &pb.PublishVideoResponse{Info: &videoInfo}, nil
}
