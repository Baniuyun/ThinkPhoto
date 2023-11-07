package logic

import (
	"Thinkphoto/server/video/rpc/code"
	"Thinkphoto/server/video/rpc/internal/svc"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除视频
func (l *DeleteVideoLogic) DeleteVideo(in *video.DeleteVideoRequest) (*video.DeleteVideoResponse, error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.TVideoModel.FindOne(l.ctx, in.VideoId)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, code.VideoNil
		}
		logx.Errorf(" l.svcCtx.TVideoModel.FindOne error: %v", err)
		return nil, err
	}
	if one.AuthorId != in.UserId {
		return nil, code.DeleteUserErr
	}
	//删除video表
	err = l.svcCtx.TVideoModel.Delete(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf(" l.svcCtx.TVideoModel.Delete error: %v", err)
		return nil, err
	}
	//删除likecount
	err = l.svcCtx.TVideoLikeCountModel.DeleteByVideoId(l.ctx, in.VideoId)
	if err != nil {
		logx.Errorf(" l.svcCtx.TVideoLikeCountModel.DeleteByVideoId error: %v", err)
		return nil, err
	}
	return &video.DeleteVideoResponse{}, nil
}
