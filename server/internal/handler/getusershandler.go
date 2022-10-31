package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"userman/server/internal/logic"
	"userman/server/internal/svc"
)

func GetUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetUsersLogic(r.Context(), svcCtx)
		resp, err := l.GetUsers()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
