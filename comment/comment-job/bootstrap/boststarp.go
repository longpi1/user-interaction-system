package bootstrap

import (
	"comment-job/libary/conf"
	"comment-job/libary/log"
	localcache "comment-job/model/dao/cache/local_cache"
	"comment-job/model/dao/cache/redis"
	"comment-job/model/dao/db"
)

func Boostrap(config *conf.WebConfig) error {
	// 启动db与cache
	client, err := db.NewClient(config)
	if err != nil && client != nil {
		log.Fatal("db run err", err)
	}
	//err = model.InitTable()
	//if err != nil {
	//	log.Fatal("初始化表失败")
	//}
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
