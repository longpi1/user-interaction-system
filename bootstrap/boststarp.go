package bootstrap

import (
	"user-interaction-system/libary/conf"
	"user-interaction-system/libary/log"
	"user-interaction-system/model/dao/cache"
	"user-interaction-system/model/dao/db"
	"user-interaction-system/model/dao/db/model"
)

func Boostrap(config *conf.Config) error {
	// 启动db与redis
	client, err := db.NewClient(config)
	if err != nil && client != nil {
		log.Fatal("db run err", err)
	}
	err = model.InitTable()
	if err != nil {
		log.Fatal("初始化表失败")
	}
	_, err = cache.NewClient(config)
	if err != nil {
		log.Fatal("cache run err", err)
	}
	return err
}
