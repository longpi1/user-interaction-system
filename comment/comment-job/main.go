package main

import (
	"comment-job/bootstrap"
	"comment-job/libary/conf"
	"comment-job/libary/log"
	"comment-job/libary/utils"
	"comment-job/model/dao/cache/redis"
)

func main() {
	config := conf.GetConfig()
	log.NewLogger(config.AppConfig.Debug, config.AppConfig.LogFilePath)

	err := bootstrap.Boostrap(config)
	if err != nil {
		log.Fatal("boostrap fail", err)
	}

	// 优雅关闭
	utils.NewHook().Close(
		// 关闭 cache
		func() {
			if redis.GetClient() != nil {
				if err := redis.GetClient().Close(); err != nil {
					log.Error("redis cache close err", err)
				}
			}
		},
	)
}
