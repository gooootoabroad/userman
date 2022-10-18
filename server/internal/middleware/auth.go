package middleware

import (
	"net/http"
	"userman/server/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.RequestURI
		logx.Infof("start check url %s auth", url)
		needAuth := true
		whiteList := config.GetConfig().WhiteList.URL
		for _, whiteURL := range whiteList {
			if whiteURL == url {
				logx.Infof("url %s is in whiteURL list, no need to check auth", url)
				needAuth = false
			}
		}

		if needAuth {
			logx.Infof("need auth")
		}
		next(w, r)
	}
}
