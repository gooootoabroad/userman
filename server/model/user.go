package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const tableName = "users"

type Role uint

var (
	undefined Role = 0
	admin     Role = 1
	sysadmin  Role = 2
	security  Role = 3
)

// 拼接密码所用的分隔符
var Separator = "."

// 错误信息
var (
	UserExist     = errors.New("username or nick name is alerady exist")
	InternalError = errors.New("occur internal error, please try again or pull an issue")
)

// UserInfo 用户信息结构体
type UserInfo struct {
	gorm.Model
	UUID     uuid.UUID  `json:"uuid"`               // 用户uuid
	Username string     `json:"username"`           // 用户名
	NickName string     `json:"nickName,omitempty"` // 用户别名
	Password string     `json:"password"`           // 密码
	Roles    Role       `json:"rolse"`              // 用户角色列表
	LockAt   *time.Time `json:"lockat,omitempty"`   // 用户锁定时间
}

// 数据库表名
func (u UserInfo) TableName() string {
	return tableName
}

func (r Role) String() (string, error) {
	if r == admin {
		return "admin", nil
	}

	if r == sysadmin {
		return "sysadmin", nil
	}

	if r == security {
		return "security", nil
	}

	return "", fmt.Errorf("undefined Role")
}

func (r Role) ToRole(role string) (Role, error) {
	if role == "admin" {
		return admin, nil
	}

	if role == "sysadmin" {
		return sysadmin, nil
	}

	if role == "security" {
		return security, nil
	}

	return undefined, fmt.Errorf("undefined Role")
}
