package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentCountLogic {
	return &UpdateCommentCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新这个视频的评论数
func (l *UpdateCommentCountLogic) UpdateCommentCount(in *pb.UpdateCommentCountReq) (*pb.UpdateCommentCountRsp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateCommentCountRsp{}, nil
}
