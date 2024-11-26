package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myprj/api/internal/logic"
	"myprj/api/internal/svc"
	"myprj/api/internal/types"
)

func ExpandHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewExpandLogic(r.Context(), svcCtx)
		resp, err := l.Expand(&req)
		if err != nil { // 如果是rpc调用的错误会有内设好的错误格式
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
