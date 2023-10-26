package logic

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/pb"
)

type GetPublishTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublishTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishTokenLogic {
	return &GetPublishTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取上传回调凭证
func (l *GetPublishTokenLogic) GetPublishToken(in *pb.GetPublishTokenRequest) (*pb.GetPublishTokenResponse, error) {
	// todo: add your logic here and delete this line
	putPolicy := storage.PutPolicy{
		Scope:            in.Bucket,
		CallbackURL:      in.CallbackURL,
		CallbackBody:     in.CallbackBody,
		CallbackBodyType: in.CallbackBodyType,
	}
	mac := qbox.NewMac(in.AccessKey, in.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	return &pb.GetPublishTokenResponse{Token: upToken}, nil
}
