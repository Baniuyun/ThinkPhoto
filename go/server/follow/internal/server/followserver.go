// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package server

import (
	"context"

	"Thinkphoto/server/follow/internal/logic"
	"Thinkphoto/server/follow/internal/svc"
	"Thinkphoto/server/follow/pb/pb"
)

type FollowServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedFollowServer
}

func NewFollowServer(svcCtx *svc.ServiceContext) *FollowServer {
	return &FollowServer{
		svcCtx: svcCtx,
	}
}

// 关注
func (s *FollowServer) Follow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {
	l := logic.NewFollowLogic(ctx, s.svcCtx)
	return l.Follow(in)
}

// 取消关注
func (s *FollowServer) UnFollow(ctx context.Context, in *pb.UnFollowRequest) (*pb.UnFollowResponse, error) {
	l := logic.NewUnFollowLogic(ctx, s.svcCtx)
	return l.UnFollow(in)
}

// 是否关注
func (s *FollowServer) IsFollow(ctx context.Context, in *pb.IsFollowRequest) (*pb.IsFollowResponse, error) {
	l := logic.NewIsFollowLogic(ctx, s.svcCtx)
	return l.IsFollow(in)
}

// 关注列表
func (s *FollowServer) FollowList(ctx context.Context, in *pb.FollowListRequest) (*pb.FollowListResponse, error) {
	l := logic.NewFollowListLogic(ctx, s.svcCtx)
	return l.FollowList(in)
}

// 粉丝列表
func (s *FollowServer) FansList(ctx context.Context, in *pb.FansListRequest) (*pb.FansListResponse, error) {
	l := logic.NewFansListLogic(ctx, s.svcCtx)
	return l.FansList(in)
}

// 关注数
func (s *FollowServer) FollowCount(ctx context.Context, in *pb.FollowCountRequest) (*pb.CountResponse, error) {
	l := logic.NewFollowCountLogic(ctx, s.svcCtx)
	return l.FollowCount(in)
}

// 粉丝数
func (s *FollowServer) FansCount(ctx context.Context, in *pb.FansCountRequest) (*pb.CountResponse, error) {
	l := logic.NewFansCountLogic(ctx, s.svcCtx)
	return l.FansCount(in)
}