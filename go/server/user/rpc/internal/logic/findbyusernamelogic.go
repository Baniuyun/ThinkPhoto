package logic

import (
	"Thinkphoto/server/user/rpc/internal/svc"
	"Thinkphoto/server/user/rpc/pb/service"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByUserNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByUserNameLogic {
	return &FindByUserNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByUserNameLogic) FindByUserName(in *service.FindByUserNameRequest) (*service.FindByUserNameResponse, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		logx.Errorf("FindByMobile mobile: %s error: %v", in.Username, err)
		return nil, err
	}
	if user == nil {
		return &service.FindByUserNameResponse{}, nil
	}

	return &service.FindByUserNameResponse{
		UserId:   user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
		Nickname: user.NickName,
		Salt:     user.Salt,
		Password: user.Password,
	}, nil
}
