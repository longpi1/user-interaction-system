package main

import (
	"relation-service/bootstrap"
	"relation-service/httpserver"
	"relation-service/libary/conf"
	"relation-service/libary/log"
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
