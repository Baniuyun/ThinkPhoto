package logic

import (
	"Thinkphoto/server/Zinc/internal/code"
	"Thinkphoto/server/Zinc/internal/svc"
	"Thinkphoto/server/Zinc/pb/zinc"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type respType struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index     string    `json:"_index"`
			Type      string    `json:"_type"`
			Id        string    `json:"_id"`
			Score     float64   `json:"_score"`
			Timestamp time.Time `json:"@timestamp"`
			Source    struct {
				Timestamp   time.Time `json:"@timestamp"`
				VideoId     int64     `json:"video_id"`
				Information string    `json:"information"`
				UserName    string    `json:"user_name"`
				UserId      int64     `json:"user_id"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *zinc.SearchRequest) (*zinc.SearchResponse, error) {
	// todo: add your logic here and delete this line
	if in.Data == "" {
		return nil, code.SearchDataEmpty
	}
	if in.SearchType == "" {
		return nil, code.SearchTypeEmpty
	}
	query := fmt.Sprintf(`{
		"search_type": "%s",
			"query":
		{
			"term": "%s"
		}
	}`, in.SearchType, in.Data)
	url := fmt.Sprintf("%s/api/%s/_search", l.svcCtx.Config.ZincSearch.Addr, l.svcCtx.Config.ZincSearch.Index)
	req, err := http.NewRequest("POST", url, strings.NewReader(query))
	if err != nil {
		logx.Errorf("zinc Search ConstructHttp err:%v", err)
		return nil, err
	}
	req.SetBasicAuth(l.svcCtx.Config.ZincSearch.UserName, l.svcCtx.Config.ZincSearch.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logx.Errorf("zinc Search HttpRequest err:%v", err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	respbody := respType{}
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &respbody)
	if err != nil {
		logx.Errorf("zinc Search json.Unmarshal err:%v", err)
		return nil, err
	}
	//fmt.Println(string(body))
	//fmt.Println(respbody)
	var index = []*zinc.Index{}
	n := len(respbody.Hits.Hits)
	for i := 0; i < n; i++ {
		index = append(index, &zinc.Index{UserId: respbody.Hits.Hits[i].Source.UserId,
			VideoId:     int64(respbody.Hits.Hits[i].Source.VideoId),
			Information: respbody.Hits.Hits[i].Source.Information,
			UserName:    respbody.Hits.Hits[i].Source.UserName,
			Id:          respbody.Hits.Hits[i].Id},
		)
	}

	return &zinc.SearchResponse{SearchResp: index}, nil
}
