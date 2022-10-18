package main

import (
	"flag"
	"fmt"

	"userman/server/global"
	"userman/server/initialize"
	"userman/server/internal/config"
	"userman/server/internal/handler"
	"userman/server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/server-api.yaml", "the config file")

func main() {
	flag.Parse()

	// 初始化配置文件
	c := config.GetConfig()
	conf.MustLoad(*configFile, c)

	// init db
	global.DB = initialize.InitDB(c.PGSQL)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(*c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
