package handler

import (
	"Thinkphoto/server/webapp/internal/logic"
	"Thinkphoto/server/webapp/internal/svc"
	"Thinkphoto/server/webapp/internal/types"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		buf := make([]byte, 1024)
		n, _ := r.Body.Read(buf)
		fmt.Println(string(rune(n)), "剑指s14")
		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
