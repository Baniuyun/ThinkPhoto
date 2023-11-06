package logic

import (
	"Thinkphoto/server/Zinc/pb/zinc"
	"Thinkphoto/server/video/rpc/code"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"
	"encoding/json"
	"strconv"

	"Thinkphoto/server/video/api/internal/svc"
	"Thinkphoto/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVideoLogic) DeleteVideo(req *types.DeleteVideoReq) (resp *types.DeleteVideoResp, err error) {
	// todo: add your logic here and delete this line
	if req.VideoId == 0 {
		return nil, code.VideoNil
	}
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	//删除video表
	_, err = l.svcCtx.VideoServer.DeleteVideo(l.ctx, &video.DeleteVideoRequest{VideoId: req.VideoId, UserId: userId})
	if err != nil {
		logx.Errorf("l.svcCtx.VideoServer.DeleteVideo err:%v", err)
		return nil, err
	}
	//删除search
	_, err = l.svcCtx.SearchServer.Delete(l.ctx, &zinc.DeleteDocRequest{Information: strconv.FormatInt(req.VideoId, 10)})
	if err != nil {
		logx.Errorf("l.svcCtx.SearchServer.Delete err:%v", err)
		return nil, err
	}
	return
}
