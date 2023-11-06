package logic

import (
	"Thinkphoto/pkg/encrypt"
	"Thinkphoto/pkg/jwt"
	"Thinkphoto/pkg/xcode"
	"Thinkphoto/server/user/rpc/user"
	"Thinkphoto/server/webapp/internal/code"
	"Thinkphoto/server/webapp/internal/svc"
	"Thinkphoto/server/webapp/internal/types"
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("start")
	req.UserName = strings.TrimSpace(req.UserName)
	if len(req.UserName) == 0 {
		return nil, code.UserNameEmpty
	}
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		return nil, code.PasswordEmpty
	}
	u, err := l.svcCtx.User.FindByUserName(l.ctx, &user.FindByUserNameRequest{Username: req.UserName})
	if err != nil {
		logx.Errorf("FindByUserName error: %v", err)
		return nil, err
	}
	if u == nil || u.UserId == 0 {
		return nil, xcode.AccessDenied
	}
	uuid := u.Salt
	Password := encrypt.EncPassword(req.Password, uuid)
	if Password != u.Password {
		return nil, code.PasswordErr
	}
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": u.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		UserId: u.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
