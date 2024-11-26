package handler

import (
	"net/http"
	"user/api/internal/logic"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		response.HttpResult(r, w, resp, err)
	}
}
