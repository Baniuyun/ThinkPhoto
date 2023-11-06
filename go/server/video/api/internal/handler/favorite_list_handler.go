package handler

import (
	"net/http"

	"Thinkphoto/server/video/api/internal/logic"
	"Thinkphoto/server/video/api/internal/svc"
	"Thinkphoto/server/video/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func favoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFavoriteListLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
