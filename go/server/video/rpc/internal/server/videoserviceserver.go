// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"rpc/internal/logic"
	"rpc/internal/svc"
	"rpc/pb"
)

type VideoServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedVideoServiceServer
}

func NewVideoServiceServer(svcCtx *svc.ServiceContext) *VideoServiceServer {
	return &VideoServiceServer{
		svcCtx: svcCtx,
	}
}

// 获取上传回调凭证
func (s *VideoServiceServer) GetPublishToken(ctx context.Context, in *pb.GetPublishTokenRequest) (*pb.GetPublishTokenResponse, error) {
	l := logic.NewGetPublishTokenLogic(ctx, s.svcCtx)
	return l.GetPublishToken(in)
}

// 获取用户上传视频列表
func (s *VideoServiceServer) GetPublishVideoList(ctx context.Context, in *pb.GetPublishVideoListRequest) (*pb.GetPublishVideoListResponse, error) {
	l := logic.NewGetPublishVideoListLogic(ctx, s.svcCtx)
	return l.GetPublishVideoList(in)
}

// 上传视频
func (s *VideoServiceServer) PublishVideo(ctx context.Context, in *pb.PublishVideoRequest) (*pb.PublishVideoResponse, error) {
	l := logic.NewPublishVideoLogic(ctx, s.svcCtx)
	return l.PublishVideo(in)
}

// 获取视频列表
func (s *VideoServiceServer) GetVideoList(ctx context.Context, in *pb.GetVideoListRequest) (*pb.GetVideoListResponse, error) {
	l := logic.NewGetVideoListLogic(ctx, s.svcCtx)
	return l.GetVideoList(in)
}

// 更新这个视频的获赞数
func (s *VideoServiceServer) UpdateFavoriteCount(ctx context.Context, in *pb.UpdateFavoriteCountReq) (*pb.UpdateFavoriteCountRsp, error) {
	l := logic.NewUpdateFavoriteCountLogic(ctx, s.svcCtx)
	return l.UpdateFavoriteCount(in)
}

// 更新这个视频的评论数
func (s *VideoServiceServer) UpdateCommentCount(ctx context.Context, in *pb.UpdateCommentCountReq) (*pb.UpdateCommentCountRsp, error) {
	l := logic.NewUpdateCommentCountLogic(ctx, s.svcCtx)
	return l.UpdateCommentCount(in)
}