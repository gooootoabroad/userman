package model

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	ExpiresTime = int64(300)
	Issuer      = "userman"
	Sign        = []byte("qgqegdbdtpgeg")
)

type Claims struct {
	UUID     uuid.UUID `gorm:"column:uuid"`     // 用户uuid
	Username string    `gorm:"column:username"` // 用户名
	Roles    string    `gorm:"column:roles"`    // 用户角色列表
	jwt.StandardClaims
}
