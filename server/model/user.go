package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const userTableName = "users"

type Role uint

var (
	Undefined Role = 0
	Admin     Role = 1
	Sysadmin  Role = 2
	Security  Role = 3
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
	UUID     uuid.UUID  `gorm:"column:uuid"`     // 用户uuid
	Username string     `gorm:"column:username"` // 用户名
	NickName string     `gorm:"column:nickname"` // 用户别名
	Password string     `gorm:"column:password"` // 密码
	Roles    Role       `gorm:"column:rolse"`    // 用户角色列表
	LockAt   *time.Time `gorm:"column:lockat"`   // 用户锁定时间
}

// 数据库表名
func (u UserInfo) TableName() string {
	return userTableName
}

func (r Role) String() (string, error) {
	if r == Admin {
		return "admin", nil
	}

	if r == Sysadmin {
		return "sysadmin", nil
	}

	if r == Security {
		return "security", nil
	}

	return "", fmt.Errorf("undefined Role")
}

func (r Role) ToRole(role string) (Role, error) {
	if role == "admin" {
		return Admin, nil
	}

	if role == "sysadmin" {
		return Sysadmin, nil
	}

	if role == "security" {
		return Security, nil
	}

	return Undefined, fmt.Errorf("undefined Role")
}
