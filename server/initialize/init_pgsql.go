package initialize

import (
	"userman/server/model"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGSQL struct {
	DNS string
}

// 初始化db
func InitDB(p PGSQL) *gorm.DB {
	if p.DNS == "" {
		panic("init db failed")
	}

	db, err := gorm.Open(postgres.Open(p.DNS), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// 初始化用户表
	initUserTable(db)
	// 初始化casbin规则表
	initCasbinPolicy(db)
	return db
}

// initUserTable 初始化用户表
func initUserTable(db *gorm.DB) {
	db.AutoMigrate(model.UserInfo{})
	// 检查超级管理员用户是否存在，不存在创建
	err := db.Where("username = ?", "admin").First(&model.UserInfo{}).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}

		password := "sha256.871659.d47174f81ece397214d27aa8ab226987db685d046a2b751f066fe0852231c543"
		adminUUID := uuid.New()
		// 没有超级管理员，创建
		admin := model.UserInfo{
			Username: "admin",
			Password: password,
			Roles:    model.Admin,
			UUID:     adminUUID,
		}

		db.Create(&admin)
	}
}

func initCasbinPolicy(db *gorm.DB) {
	db.AutoMigrate(model.CasbinRule{})
	err := db.Where("V0 = ?", "admin").First(&model.CasbinRule{}).Error
	if err != nil {
		// 没有超级管理员访问的接口，一般是有问题的，重新写入默认规则
		policy := []model.CasbinRule{
			{Ptype: "p", V0: "admin", V1: "/user", V2: "GET"},
			{Ptype: "p", V0: "admin", V1: "/user", V2: "POST"},
			{Ptype: "p", V0: "admin", V1: "/user", V2: "DELETE"},
			{Ptype: "p", V0: "admin", V1: "/users", V2: "GET"},
			{Ptype: "p", V0: "admin", V1: "/login", V2: "POST"},
			{Ptype: "p", V0: "sysadmin", V1: "/user", V2: "GET"},
			{Ptype: "p", V0: "sysadmin", V1: "/user", V2: "POST"},
			{Ptype: "p", V0: "sysadmin", V1: "/user", V2: "DELETE"},
			{Ptype: "p", V0: "sysadmin", V1: "/login", V2: "POST"},
			{Ptype: "p", V0: "security", V1: "/login", V2: "POST"},
		}

		db.Create(&policy)
	}
}
