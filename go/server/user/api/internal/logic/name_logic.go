package logic

import (
	"context"

	"Thinkphoto/server/user/api/internal/svc"
	"Thinkphoto/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NameLogic {
	return &NameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NameLogic) Name(req *types.NameReq) (resp *types.NameResp, err error) {
	// todo: add your logic here and delete this line
	//l.svcCtx.UserServer
	return
}
