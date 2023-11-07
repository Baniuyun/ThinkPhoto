package logic

import (
	"Thinkphoto/server/user/rpc/internal/svc"
	"Thinkphoto/server/user/rpc/pb/service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *service.FindByIdRequest) (*service.FindByIdResponse, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.TUserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("FindById userId: %d error: %v", in.UserId, err)
		return nil, err
	}
	return &service.FindByIdResponse{
		UserId:   user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
		Nickname: user.NickName,
		Password: user.Password,
		Salt:     user.Salt,
	}, nil
}
