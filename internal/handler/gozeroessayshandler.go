package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_essays/internal/logic"
	"go_zero_essays/internal/svc"
	"go_zero_essays/internal/types"
)

func Go_zero_essaysHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGo_zero_essaysLogic(r.Context(), svcCtx)
		resp, err := l.Go_zero_essays(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
