package logic

import (
	"Thinkphoto/server/Zinc/internal/code"
	"Thinkphoto/server/Zinc/internal/svc"
	"Thinkphoto/server/Zinc/pb/zinc"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *zinc.DeleteDocRequest) (*zinc.Response, error) {
	// todo: add your logic here and delete this line
	if in.Information == "" {
		return nil, code.DeleteInformationEmpty
	}
	ls := &SearchLogic{l.ctx, l.svcCtx, l.Logger}
	searchin := &zinc.SearchRequest{SearchType: "match", Data: in.Information}
	searchresp, err := ls.Search(searchin)
	if err != nil {
		logx.Errorf("zinc delete use search err:%v", err)
		return nil, err
	}
	id := searchresp.SearchResp[0].Id
	url := fmt.Sprintf("%s/api/%s/_doc/%s", l.svcCtx.Config.ZincSearch.Addr, l.svcCtx.Config.ZincSearch.Index, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		logx.Errorf("zinc delete ConstructHttp err:%v", err)
		return nil, err
	}
	req.SetBasicAuth(l.svcCtx.Config.ZincSearch.UserName, l.svcCtx.Config.ZincSearch.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logx.Errorf("zinc delete HttpRequest err:%v", err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	return &zinc.Response{Msg: string(body)}, nil
}
