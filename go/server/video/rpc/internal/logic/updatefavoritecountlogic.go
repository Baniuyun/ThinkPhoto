package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

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

// 更新这个视频的获赞数
func (l *UpdateFavoriteCountLogic) UpdateFavoriteCount(in *pb.UpdateFavoriteCountReq) (*pb.UpdateFavoriteCountRsp, error) {
	videoId := in.VideoId
	favoriteStatus := in.FavoriteStatus

	if favoriteStatus == 0 {
		return nil, nil
	}
	_, err := l.svcCtx.TVideoModel.UpdateFavoriteCount(context.Background(), videoId)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateFavoriteCountRsp{CommonRsp: &pb.CommonResponse{Code: 200, Msg: "update favorite_count success"}}, nil
}
