package main

import (
	"github.com/longpi1/user-interaction-system/comment/comment-job/bootstrap"
	"github.com/longpi1/user-interaction-system/comment/comment-job/httpserver"
	"github.com/longpi1/user-interaction-system/comment/comment-job/job"
	"github.com/longpi1/user-interaction-system/comment/comment-job/libary/conf"
	"github.com/longpi1/user-interaction-system/comment/comment-job/model/dao/cache/redis"

	"github.com/longpi1/gopkg/libary/log"
	"github.com/longpi1/gopkg/libary/utils"
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

	// 启动任务
	job.Start()

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
