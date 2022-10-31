package utils

import (
	"context"
	"userman/server/global"
	"userman/server/model"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

func Casbin(ctx context.Context) (*casbin.Enforcer, error) {
	logger := logx.WithContext(ctx)
	casbinRule := &model.CasbinRule{}
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(global.DB, casbinRule, casbinRule.TableName())
	if err != nil {
		logger.Errorf("create gorm adapter failed, err: %v", err)
		return nil, err
	}

	e, err := casbin.NewEnforcer("etc/model.conf", adapter)
	if err != nil {
		logger.Errorf("create casbin enforcer failed, err: %v", err)
		return nil, err
	}

	e.LoadPolicy()
	return e, nil
}
