package bootstrap

import (
	"fmt"
	"model-api/libary/conf"
	"model-api/libary/log"
	"model-api/model/dao/cache"
	"model-api/model/dao/db"
	"model-api/model/dao/db/model"
)

func Boostrap() error {
	var dbConfig conf.DBConf
	// 构建DSN字符串
	dbConfig.Dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		"root",
		"123456",
		"127.0.0.1:3306",
		"model_api",
		true,
		"Local")
	dbConfig.Type = "mysql"
	// 启动db与redis
	client, err := db.NewClient(dbConfig)
	if err != nil && client != nil {
		log.Fatal("db run err", err)
	}
	err = model.InitTable()
	if err != nil {
		log.Fatal("初始化表失败")
	}
	var redisConfig conf.RedisConf
	_, err = cache.NewClient(redisConfig)
	if err != nil {
		log.Fatal("cache run err", err)
	}
	return err
}
