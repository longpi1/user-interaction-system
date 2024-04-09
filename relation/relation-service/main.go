package main

import (
	"relation-service/bootstrap"
	"relation-service/httpserver"
	"relation-service/libary/conf"
	"relation-service/libary/log"
	"relation-service/libary/utils"
	localcache "relation-service/model/dao/cache/local_cache"
	"relation-service/model/dao/cache/redis"
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

	// 优雅关闭
	utils.NewHook().Close(
		// 关闭 cache
		func() {
			if localcache.GetClient() != nil {
				if err := localcache.GetClient().Close(); err != nil {
					log.Error("local cache close err", err)
				}
			}
			if redis.GetClient() != nil {
				if err := redis.GetClient().Close(); err != nil {
					log.Error("redis cache close err", err)
				}
			}
		},
	)
}
