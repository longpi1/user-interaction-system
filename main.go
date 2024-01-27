package main

import (
	"user-interaction-system/bootstrap"
	"user-interaction-system/httpserver"
	"user-interaction-system/libary/conf"
	"user-interaction-system/libary/log"
)

func main() {
	config := conf.GetConfig()
	log.NewLogger(config.AppConfig.Debug, config.AppConfig.LogFilePath)

	err := bootstrap.Boostrap(config)
	if err != nil {
		log.Fatal("boostrap fail", err)
	}
	// 启动http服务路由
	httpserver.SetRouter(config.AppConfig.Port)
}
