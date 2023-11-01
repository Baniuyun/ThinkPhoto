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

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *zinc.UpdateDocRequest) (*zinc.Response, error) {
	// todo: add your logic here and delete this line
	ls := &SearchLogic{l.ctx, l.svcCtx, l.Logger}
	searchin := &zinc.SearchRequest{SearchType: "match", Data: in.VideoId}
	searchresp, err := ls.Search(searchin)
	if err != nil {
		return nil, err
	}
	id := searchresp.SearchResp[0].Id
	url := fmt.Sprintf("%s/api/%s/_update/%s", l.svcCtx.Config.ZincSearch.Addr, l.svcCtx.Config.ZincSearch.Index, id)
	requestbody := fmt.Sprintf(`{"video_id":"%s",
										"information":"%s",
										"user_name":"%s",
										"user_id":"%s"}`, in.Data.VideoId, in.Data.Information, in.Data.UserName, in.Data.UserId)
	req, err := http.NewRequest("POST", url, strings.NewReader(requestbody))
	if err != nil {
		logx.Errorf("zinc Update ConstructHttp err:%v", err)
		return nil, err
	}
	req.SetBasicAuth(l.svcCtx.Config.ZincSearch.UserName, l.svcCtx.Config.ZincSearch.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logx.Errorf("zinc Update HttpRequest err:%v", err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	return &zinc.Response{Msg: string(body)}, nil
}
