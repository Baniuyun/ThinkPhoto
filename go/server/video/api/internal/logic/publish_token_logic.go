package logic

import (
	"Thinkphoto/pkg/xcode"
	"Thinkphoto/server/video/api/internal/svc"
	"Thinkphoto/server/video/api/internal/types"
	"Thinkphoto/server/video/rpc/pb/video"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishTokenLogic {
	return &PublishTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishTokenLogic) PublishToken(req *types.PublishTokenReq) (*types.PublishTokenResp, error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	if userId == 0 {
		return nil, xcode.NoLogin
	}
	callBackbody := `{"key":"$(key)",
						"hash":"$(etag)",
						"tag":"$(x:tag)",
						"information":"$(x:information)",
						"userId":"$(x:userId)",
						"userName":"$(x:userName)"
}`
	token, err := l.svcCtx.VideoServer.GetPublishToken(l.ctx, &video.GetPublishTokenRequest{
		AccessKey:        l.svcCtx.Config.Qiniu.AccessKeyId,
		SecretKey:        l.svcCtx.Config.Qiniu.SecretAccessKey,
		Bucket:           l.svcCtx.Config.Qiniu.VideoBuckets,
		CallbackURL:      "42.194.236.154:8888",
		CallbackBody:     callBackbody,
		CallbackBodyType: "application/json",
	})
	fmt.Println(token.Token)
	if err != nil {
		logx.Errorf("l.svcCtx.VideoServer.GetPublishToken err %v", err)
		return nil, err
	}
	return &types.PublishTokenResp{
		Token: token.Token}, nil
}
