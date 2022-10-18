package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	ExpiresTime = int64(time.Minute * 5)
	Issuer      = "userman"
	Sign        = []byte("qgqegdbdtpgeg")
)

type Claims struct {
	UUID     uuid.UUID `json:"uuid"`     // 用户uuid
	Username string    `json:"username"` // 用户名
	Roles    string    `json:"roles"`    // 用户角色列表
	jwt.StandardClaims
}
