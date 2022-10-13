package initialize

import (
	"userman/server/model"

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
		panic("init db failed")
	}

	db.AutoMigrate(model.UserInfo{})
	return db
}
