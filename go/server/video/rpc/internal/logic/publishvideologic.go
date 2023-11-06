package logic

import (
	"Thinkphoto/server/video/rpc/internal/model"
	"Thinkphoto/server/video/rpc/internal/svc"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"
	"database/sql"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type callBackBody struct {
	key         string `json:"key"`
	hash        string `json:"hash"`
	VideoId     int64  `json:"videoId"`
	Tag         int64  `json:"tag"`
	Information string `json:"information"`
	UserId      int64  `json:"userId"`
	UserName    string `json:"userName"`
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传视频
func (l *PublishVideoLogic) PublishVideo(in *video.PublishVideoRequest) (*video.PublishVideoResponse, error) {
	// todo: add your logic here and delete this line
	//var body callBackBody
	//if err := json.Unmarshal([]byte(in, &body); err != nil {
	//	logx.Errorf("Failed to unmarshal JSON: %v", err)
	//	return nil, err
	//}

	//qiniu := l.svcCtx.Config.Qiniu
	//videoInfo.PlayUrl = qiniu.Addr + qiniu.VideoPath + ".mp4"
	//videoInfo.CoverUrl = qiniu.Addr + qiniu.VideoPath + ".jpg"

	mac := qbox.NewMac(l.svcCtx.Config.Qiniu.AccessKeyId, l.svcCtx.Config.Qiniu.SecretAccessKey)
	domain := "www.s32fnznta.hd-bkt.clouddn.com"
	key := in.Key + ".mp4"
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)

	tVideo := model.TVideo{
		AuthorId:   in.UserId,
		AuthorName: in.Username,
		PlayUrl:    privateAccessURL,
		CoverUrl: sql.NullString{
			String: "",
			Valid:  false,
		},
		Avatar:        in.Avatar,
		FavoriteCount: 0,
		CommentCount:  0,
		Information:   in.Information,
		Tag:           in.Tag,
		Time:          time.Now().Unix(),
	}

	res, err := l.svcCtx.TVideoModel.Insert(context.Background(), &tVideo)
	if err != nil {
		logx.Errorf("l.svcCtx.TVideoModel insert fail: %v", err)
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		logx.Errorf("res.LastInsertId fail: %v", err)
		return nil, err
	}
	//videoInfo.VideoId = id

	return &video.PublishVideoResponse{Info: &video.VideoInfo{
		VideoId: id,
		//AuthorId:      ,
		PlayUrl: privateAccessURL,
		//CoverUrl:      "",
		//IsFavorite:    false,
	}}, nil
}
