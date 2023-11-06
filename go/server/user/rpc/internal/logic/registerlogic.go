package logic

import (
	"Thinkphoto/pkg/encrypt"
	"Thinkphoto/server/user/rpc/internal/code"
	"Thinkphoto/server/user/rpc/internal/model"
	"Thinkphoto/server/user/rpc/internal/svc"
	"Thinkphoto/server/user/rpc/pb/service"
	"context"
	uuid3 "github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	if in.Password == "" {
		return nil, code.RegisterPasswordEmpty
	}
	if in.Username == "" {
		return nil, code.RegisterNameEmpty
	}
	_, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return nil, code.RegisterNameRepeat
	}
	if err != sqlx.ErrNotFound {
		logx.Errorf("FindOneByUsername req: %v error: %v", in, err)
		return nil, err
	}
	uuid := uuid3.NewString()
	newPassword := encrypt.EncPassword(in.Password, uuid)
	ret, err := l.svcCtx.TUserModel.Insert(l.ctx, &model.TUser{
		Username:   in.Username,
		Salt:       uuid,
		Password:   newPassword,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		NickName:   "",
		Avatar:     "",
	})
	if err != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, err
	}
	userId, err := ret.LastInsertId()
	if err != nil {
		logx.Errorf("LastInsertId error: %v", err)
		return nil, err
	}
	return &service.RegisterResponse{UserId: userId}, nil
}
