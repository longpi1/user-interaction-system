package bootstrap

import (
	"github.com/longpi1/user-interaction-system/comment-job/libary/conf"
	"github.com/longpi1/user-interaction-system/comment-job/model/dao/cache/redis"
	"github.com/longpi1/user-interaction-system/comment-job/model/dao/db"
	"github.com/longpi1/user-interaction-system/comment-job/model/dao/db/model"

	"github.com/longpi1/gopkg/libary/log"
)

func Boostrap(config *conf.WebConfig) error {
	// 启动db与cache
	client, err := db.NewClient(config)
	if err != nil && client != nil {
		log.Fatal("db run err", err)
	}
	err = model.InitTable()
	if err != nil {
		log.Fatal("初始化表失败")
	}
	_, err = redis.NewClient(config)
	if err != nil {
		log.Fatal("redis cache run err", err)
	}

	return err
}
