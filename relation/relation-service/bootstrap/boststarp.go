package bootstrap

import (
	"relation-service/libary/conf"
	"relation-service/libary/log"
	localcache "relation-service/model/dao/cache/local_cache"
	"relation-service/model/dao/cache/redis"
	"relation-service/model/dao/db"
	"relation-service/model/dao/db/model"
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
