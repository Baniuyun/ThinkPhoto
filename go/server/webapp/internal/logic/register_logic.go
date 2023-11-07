package logic

import (
	"Thinkphoto/pkg/jwt"
	"Thinkphoto/server/user/rpc/user"
	"Thinkphoto/server/webapp/internal/svc"
	"Thinkphoto/server/webapp/internal/types"
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	req.UserName = strings.TrimSpace(req.UserName)
	req.Password = strings.TrimSpace(req.Password)
	//uuid := uuid3.NewString()
	//req.Password = encrypt.EncPassword(req.Password, uuid)
	//req.VerificationCode = strings.TrimSpace(req.VerificationCode)

	//u, err := l.svcCtx.User.FindByUserName(l.ctx, &user.FindByUserNameRequest{Username: req.UserName})
	//if err != nil {
	//	logx.Errorf("FindByUserName error: %v", err)
	//	return nil, err
	//}
	//if u != nil && u.UserId > 0 {
	//	return nil, code.UserNameHasRegistered
	//}
	regRet, err := l.svcCtx.User.Register(l.ctx, &user.RegisterRequest{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("Register error: %v", err)
		return nil, err
	}
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": regRet.UserId,
		},
	})
	if err != nil {
		logx.Errorf("BuildTokens error: %v", err)
		return nil, err
	}
	//_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.BizRedis)
	return &types.RegisterResponse{
		UserId: regRet.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
	return
}
