package logic

import (
	"context"

	"Thinkphoto/server/user/api/internal/svc"
	"Thinkphoto/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishAvatarLogic {
	return &PublishAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishAvatarLogic) PublishAvatar(req *types.PublishAvatarReq) (resp *types.PublishAvatarResp, err error) {
	// todo: add your logic here and delete this line

	return
}
