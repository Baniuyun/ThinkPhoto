package handler

import (
	"net/http"

	"Thinkphoto/server/user/api/internal/logic"
	"Thinkphoto/server/user/api/internal/svc"
	"Thinkphoto/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getVideoInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetVideoInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetVideoInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
