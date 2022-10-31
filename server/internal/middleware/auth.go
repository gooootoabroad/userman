package middleware

import (
	"net/http"
	"time"
	"userman/server/internal/config"
	"userman/server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logx.WithContext(r.Context())
		url := r.RequestURI
		logger.Infof("start check url %s auth", url)
		needAuth := false
		whiteList := config.Get().WhiteList.URL
		for _, whiteURL := range whiteList {
			if whiteURL == url {
				logger.Infof("url %s is in whiteURL list, no need to check auth", url)
				needAuth = false
			}
		}

		if needAuth {
			logger.Infof("url %s need auth", url)
			auth := r.Header.Get("Authorization")
			if auth == "" {
				logger.Errorf("url %s not find auth info", url)
				http.Error(w, "Need Auth", http.StatusUnauthorized)
				return
			}

			// 校验jwt
			claims, err := utils.DecodeToken(r.Context(), auth)
			if err != nil {
				logger.Errorf("verify jwt %s failed, err: %v", auth, err)
				http.Error(w, "Token is not invalid", http.StatusUnauthorized)
				return
			}

			// 检查是否过期了
			logger.Infof("now %v e %v", time.Now().Unix(), claims.ExpiresAt)
			if time.Now().Unix() > claims.ExpiresAt {
				logger.Errorf("jwt %s is expired, expire at %d", auth, claims.ExpiresAt)
				http.Error(w, "Token is expired", http.StatusUnauthorized)
				return
			}
			e, err := utils.Casbin(r.Context())
			if err != nil {
				logger.Errorf("get casbin adapter failed, err: %v", err)
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}

			ok, reason, err := e.EnforceEx(claims.Username, url, r.Method)
			logger.Infof("user %s access url %s result: %b reason: %v", claims.Username, url, ok, reason)
			if !ok {
				logger.Errorf("user %s no auth to access url %s,", claims.Username, url)
				http.Error(w, "No premission", http.StatusBadRequest)
				return
			}
		}
		next(w, r)
	}
}
