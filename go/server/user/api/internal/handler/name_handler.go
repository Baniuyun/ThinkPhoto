package handler

import (
	"net/http"

	"Thinkphoto/server/user/api/internal/logic"
	"Thinkphoto/server/user/api/internal/svc"
	"Thinkphoto/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func nameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewNameLogic(r.Context(), svcCtx)
		resp, err := l.Name(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
