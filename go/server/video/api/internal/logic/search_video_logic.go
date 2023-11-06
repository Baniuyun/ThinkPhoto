package logic

import (
	"Thinkphoto/server/Zinc/zincsearch"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"

	"Thinkphoto/server/video/api/internal/svc"
	"Thinkphoto/server/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchVideoLogic {
	return &SearchVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchVideoLogic) SearchVideo(req *types.SearchReq) (*types.SearchResp, error) {
	// todo: add your logic here and delete this line
	search, err := l.svcCtx.SearchServer.Search(l.ctx, &zincsearch.SearchRequest{Data: req.Word, SearchType: "match"})
	if err != nil {
		logx.Errorf("l.svcCtx.SearchServer.Search err: %v", err)
		return nil, err
	}
	videolist := []types.VideoSimple{}
	n := len(search.SearchResp)
	for i := 0; i < n; i++ {
		id, err := l.svcCtx.VideoServer.GetVideoInfoBYVideoId(l.ctx, &video.GetVideoInfoReq{
			VideoId: search.SearchResp[i].VideoId,
			UserId:  search.SearchResp[i].UserId,
		})
		if err != nil {
			logx.Errorf("l.svcCtx.VideoServer.GetVideoInfoBYVideoId err : %v", err)
			return nil, err
		}
		videolist = append(videolist, types.VideoSimple{
			VideoId:       id.Info.VideoId,
			CoverURL:      id.Info.PlayUrl,
			FavoriteCount: id.Info.FavoriteCount,
			Information:   id.Info.Information,
		})
	}
	return &types.SearchResp{VideoList: videolist}, nil
}
