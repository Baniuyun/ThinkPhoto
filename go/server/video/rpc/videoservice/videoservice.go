// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videoservice

import (
	"Thinkphoto/server/video/rpc/pb/video"
	"context"



	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonResponse               = video.CommonResponse
	DeleteVideoRequest           = video.DeleteVideoRequest
	DeleteVideoResponse          = video.DeleteVideoResponse
	GetFavoriteVideoListRequest  = video.GetFavoriteVideoListRequest
	GetFavoriteVideoListResponse = video.GetFavoriteVideoListResponse
	GetPublishTokenRequest       = video.GetPublishTokenRequest
	GetPublishTokenResponse      = video.GetPublishTokenResponse
	GetPublishVideoListRequest   = video.GetPublishVideoListRequest
	GetPublishVideoListResponse  = video.GetPublishVideoListResponse
	GetVideoInfoReq              = video.GetVideoInfoReq
	GetVideoInfoResp             = video.GetVideoInfoResp
	GetVideoListRequest          = video.GetVideoListRequest
	GetVideoListResponse         = video.GetVideoListResponse
	PublishVideoRequest          = video.PublishVideoRequest
	PublishVideoResponse         = video.PublishVideoResponse
	UpdateFavoriteCountReq       = video.UpdateFavoriteCountReq
	UpdateFavoriteCountRsp       = video.UpdateFavoriteCountRsp
	VideoInfo                    = video.VideoInfo
	VideoSimple                  = video.VideoSimple

	VideoService interface {
		// 获取上传回调凭证
		GetPublishToken(ctx context.Context, in *GetPublishTokenRequest, opts ...grpc.CallOption) (*GetPublishTokenResponse, error)
		// 获取用户上传视频列表
		GetPublishVideoList(ctx context.Context, in *GetPublishVideoListRequest, opts ...grpc.CallOption) (*GetPublishVideoListResponse, error)
		// 获取用户喜欢视频列表
		GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListResponse, error)
		// 上传视频
		PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error)
		// 删除视频
		DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error)
		// 获取视频列表
		GetVideoList(ctx context.Context, in *GetVideoListRequest, opts ...grpc.CallOption) (*GetVideoListResponse, error)
		// 点赞并更新这个视频的获赞数
		UpdateFavoriteCount(ctx context.Context, in *UpdateFavoriteCountReq, opts ...grpc.CallOption) (*UpdateFavoriteCountRsp, error)
		// 根据视频id获取视频信息
		GetVideoInfoBYVideoId(ctx context.Context, in *GetVideoInfoReq, opts ...grpc.CallOption) (*GetVideoInfoResp, error)
	}

	defaultVideoService struct {
		cli zrpc.Client
	}
)

func NewVideoService(cli zrpc.Client) VideoService {
	return &defaultVideoService{
		cli: cli,
	}
}

// 获取上传回调凭证
func (m *defaultVideoService) GetPublishToken(ctx context.Context, in *GetPublishTokenRequest, opts ...grpc.CallOption) (*GetPublishTokenResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetPublishToken(ctx, in, opts...)
}

// 获取用户上传视频列表
func (m *defaultVideoService) GetPublishVideoList(ctx context.Context, in *GetPublishVideoListRequest, opts ...grpc.CallOption) (*GetPublishVideoListResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetPublishVideoList(ctx, in, opts...)
}

// 获取用户喜欢视频列表
func (m *defaultVideoService) GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetFavoriteVideoList(ctx, in, opts...)
}

// 上传视频
func (m *defaultVideoService) PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.PublishVideo(ctx, in, opts...)
}

// 删除视频
func (m *defaultVideoService) DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.DeleteVideo(ctx, in, opts...)
}

// 获取视频列表
func (m *defaultVideoService) GetVideoList(ctx context.Context, in *GetVideoListRequest, opts ...grpc.CallOption) (*GetVideoListResponse, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetVideoList(ctx, in, opts...)
}

// 点赞并更新这个视频的获赞数
func (m *defaultVideoService) UpdateFavoriteCount(ctx context.Context, in *UpdateFavoriteCountReq, opts ...grpc.CallOption) (*UpdateFavoriteCountRsp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.UpdateFavoriteCount(ctx, in, opts...)
}

// 根据视频id获取视频信息
func (m *defaultVideoService) GetVideoInfoBYVideoId(ctx context.Context, in *GetVideoInfoReq, opts ...grpc.CallOption) (*GetVideoInfoResp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetVideoInfoBYVideoId(ctx, in, opts...)
}
