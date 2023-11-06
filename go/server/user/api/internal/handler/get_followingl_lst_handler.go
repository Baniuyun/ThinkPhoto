package handler

import (
	"net/http"

	"Thinkphoto/server/user/api/internal/logic"
	"Thinkphoto/server/user/api/internal/svc"
	"Thinkphoto/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getFollowinglLstHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowinglistReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetFollowinglLstLogic(r.Context(), svcCtx)
		resp, err := l.GetFollowinglLst(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
