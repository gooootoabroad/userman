package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const tableName = "users"

// UserInfo 用户信息结构体
type UserInfo struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid"`               // 用户uuid
	Username string    `json:"username"`           // 用户名
	NickName string    `json:"nickname,omitempty"` // 用户别名
	Password string    `json:"password"`           // 密码
	Roles    uint      `json:"rolse"`              // 用户角色列表
	LockAt   time.Time `json:"lockat,omitempty"`   // 用户锁定时间
}

// 数据库表名
func (u UserInfo) TableName() string {
	return tableName
}
