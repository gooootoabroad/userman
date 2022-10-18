package model

import "github.com/google/uuid"

type JWTInfo struct {
	UUID     uuid.UUID `json:"uuid"`     // 用户uuid
	ID       uint      `json:"id"`       // 数据库主键
	Username string    `json:"username"` //用户名
	Roles    uint      `json:"roles"`    // 用户角色列表
}
