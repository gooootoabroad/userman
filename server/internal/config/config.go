package config

import (
	"userman/server/initialize"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var c = Config{}

type Config struct {
	rest.RestConf
	initialize.PGSQL
	RSA
	WhiteList WhiteURL
}

type RSA struct {
	PublicKeyFile  string `json:""`
	PrivateKeyFile string `json:""`
}

type WhiteURL struct {
	URL []string
}

func InitWithFatal(path string) Config {
	conf.MustLoad(path, &c)
	return c
}

func Get() Config {
	return c
}
