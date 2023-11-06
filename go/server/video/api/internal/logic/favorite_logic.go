package logic

import (
	"Thinkphoto/pkg/xcode"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"
	"encoding/json"

	"Thinkphoto/server/video/api/internal/svc"
	"Thinkphoto/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteLogic) Favorite(req *types.FavoriteReq) (*types.FavoriteResp, error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	if userId == 0 {
		return nil, xcode.NoLogin
	}
	_, err = l.svcCtx.VideoServer.UpdateFavoriteCount(l.ctx, &video.UpdateFavoriteCountReq{
		VideoId: req.VideoId,
		UserId:  req.UserId,
		Status:  req.Status})
	if err != nil {
		return nil, err
	}
	return &types.FavoriteResp{}, nil
}
