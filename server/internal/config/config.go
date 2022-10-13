package config

import (
	"userman/server/initialize"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	initialize.PGSQL
}
