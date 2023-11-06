package logic

import (
	"Thinkphoto/server/video/rpc/internal/model"
	"Thinkphoto/server/video/rpc/internal/svc"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFavoriteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFavoriteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFavoriteCountLogic {
	return &UpdateFavoriteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞并更新这个视频的获赞数
func (l *UpdateFavoriteCountLogic) UpdateFavoriteCount(in *video.UpdateFavoriteCountReq) (*video.UpdateFavoriteCountRsp, error) {
	// todo: add your logic here and delete this line
	//更新t_video标
	_, err := l.svcCtx.TVideoModel.UpdateFavoriteCountByStatus(l.ctx, in.Status, in.VideoId)
	if err != nil {
		logx.Errorf("l.svcCtx.TVideoModel.UpdateFavoriteCountByStatus err : %v", err)
		return nil, err
	}
	//更新like_record表
	if in.Status == 0 {
		_, err = l.svcCtx.TVideoLikeRecordModel.Insert(l.ctx, &model.LikeRecord{LinkingId: in.VideoId, LikerId: in.UserId})
		if err != nil {
			logx.Errorf("l.svcCtx.TVideoLikeRecordModel.Insert err : %v", err)
			return nil, err
		}
	} else {
		err = l.svcCtx.TVideoLikeRecordModel.DeleteByLinkIdAndLikerId(l.ctx, in.VideoId, in.UserId)
		if err != nil {
			logx.Errorf("l.svcCtx.TVideoLikeRecordModel.DeleteByLinkIdAndLikerId err : %v", err)
			return nil, err
		}
	}
	//更新like_count表
	_, err = l.svcCtx.TVideoLikeCountModel.UpdateFavoriteCount(l.ctx, in.Status, in.VideoId)
	if err != nil {
		logx.Errorf("l.svcCtx.TVideoLikeCountModel.UpdateFavoriteCount err : %v", err)
		return nil, err
	}
	return &video.UpdateFavoriteCountRsp{}, nil
}
