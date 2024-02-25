package bootstrap

import (
	"user-interaction-system/libary/conf"
	"user-interaction-system/libary/log"
	localcache "user-interaction-system/model/dao/cache/local_cache"
	"user-interaction-system/model/dao/cache/redis"
	"user-interaction-system/model/dao/db"
	"user-interaction-system/model/dao/db/model"
)

func Boostrap(config *conf.Config) error {
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
