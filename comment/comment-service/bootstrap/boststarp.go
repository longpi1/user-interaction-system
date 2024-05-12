package bootstrap

import (
	"github.com/longpi1/user-interaction-system/comment/comment-service/libary/conf"
	localcache "github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/cache/local_cache"
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/cache/redis"
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db"
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db/model"

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
	_, err = localcache.NewClient(config)
	if err != nil {
		log.Fatal("local cache run err", err)
	}

	return err
}
