package logic

import (
	"Thinkphoto/server/Zinc/internal/svc"
	"Thinkphoto/server/Zinc/pb/zinc"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type DocLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DocLogic {
	return &DocLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DocLogic) Doc(in *zinc.Doc) (*zinc.Response, error) {
	// todo: add your logic here and delete this line
	requestbody := fmt.Sprintf(`{"video_id":"%s",
										"information":"%s",
										"user_name":"%s",
										"user_id":"%s"}`, in.VideoId, in.Information, in.UserName, in.UserId)
	url := fmt.Sprintf("%s/api/%s/_doc", l.svcCtx.Config.ZincSearch.Addr, l.svcCtx.Config.ZincSearch.Index)
	req, err := http.NewRequest("POST", url, strings.NewReader(requestbody))
	if err != nil {
		logx.Errorf("zinc Doc  ConstructHttp err:%v", err)
		return nil, err
	}
	req.SetBasicAuth(l.svcCtx.Config.ZincSearch.UserName, l.svcCtx.Config.ZincSearch.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logx.Errorf("zinc Doc HttpRequest err:%v", err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	return &zinc.Response{Msg: string(body)}, nil
}
