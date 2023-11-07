package handler

import (
	"Thinkphoto/server/webapp/internal/logic"
	"Thinkphoto/server/webapp/internal/svc"
	"Thinkphoto/server/webapp/internal/types"
	"fmt"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		body, err := io.ReadAll(r.Body)
		fmt.Println(string(body), "剑指s14")
		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
