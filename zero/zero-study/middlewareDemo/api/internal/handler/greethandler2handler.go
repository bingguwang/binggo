package handler

import (
	"net/http"

	"binggo/zero/zero-study/middlewareDemo/api/internal/logic"
	"binggo/zero/zero-study/middlewareDemo/api/internal/svc"
	"binggo/zero/zero-study/middlewareDemo/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GreetHandler2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGreetHandler2Logic(r.Context(), svcCtx)
		resp, err := l.GreetHandler2(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}