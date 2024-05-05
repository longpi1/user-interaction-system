package main

import (
	"comment-job/bootstrap"
	"comment-job/libary/conf"
	"comment-job/model/dao/cache/redis"
	"context"

	"github.com/longpi1/gopkg/libary/queue"
	"github.com/longpi1/gopkg/libary/utils"

	"github.com/longpi1/gopkg/libary/log"
)

func main() {
	config := conf.GetConfig()
	log.NewLogger(config.AppConfig.Debug, config.AppConfig.LogFilePath)

	err := bootstrap.Boostrap(config)
	if err != nil {
		log.Fatal("boostrap fail", err)
	}

	// 启动消费者对队列数据进行处理
	queue.StartConsumersListener(context.Background(), config.QueueConfig.)

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
