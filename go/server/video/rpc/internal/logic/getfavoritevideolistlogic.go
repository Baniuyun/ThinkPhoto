package logic

import (
	"Thinkphoto/server/video/rpc/code"
	"Thinkphoto/server/video/rpc/internal/svc"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteVideoListLogic {
	return &GetFavoriteVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户喜欢视频列表
func (l *GetFavoriteVideoListLogic) GetFavoriteVideoList(in *video.GetFavoriteVideoListRequest) (*video.GetFavoriteVideoListResponse, error) {
	// todo: add your logic here and delete this line
	id, err := l.svcCtx.TVideoLikeRecordModel.FindLikeListById(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, code.UserIdEmpty
		}
		logx.Errorf("l.svcCtx.TVideoLikeRecordModel.FindLikeListById err%v", err)
		return nil, err
	}
	var list []*video.VideoSimple
	for _, r := range id {
		one, err := l.svcCtx.TVideoModel.FindOne(l.ctx, r.LinkingId)
		if err != nil {
			if errors.Is(err, sqlx.ErrNotFound) {
				return nil, code.VideoNil
			}
			logx.Errorf("l.svcCtx.TVideoModel.FindOne err%v", err)
			return nil, err
		}
		list = append(list, &video.VideoSimple{
			VideoId:       one.Id,
			CoverUrl:      one.PlayUrl,
			Information:   one.Information,
			Time:          one.Time,
			FavoriteCount: one.FavoriteCount,
		})
	}
	return &video.GetFavoriteVideoListResponse{VideoList: list}, nil
}
