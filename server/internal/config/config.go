package config

import (
	"userman/server/initialize"

	"github.com/zeromicro/go-zero/rest"
)

var c = &Config{}

type Config struct {
	rest.RestConf
	initialize.PGSQL
	RSA
	WhiteList WhiteList
}

type RSA struct {
	PublicKeyFile  string `json:""`
	PrivateKeyFile string `json:""`
}

type WhiteList struct {
	URL []string
}

func GetConfig() *Config {
	return c
}
