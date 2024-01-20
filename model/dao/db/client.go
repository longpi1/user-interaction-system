package db

import (
	"fmt"
	"gorm.io/gorm"
	"model-api/libary/conf"
	"model-api/libary/log"
	"sync"
)

var db = &DB{}
var once = sync.Once{}

type DB struct {
	client *gorm.DB
}

func GetClient() *gorm.DB {
	if db == nil || db.client == nil {
		once.Do(func() {
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
			client, err := NewClient(dbConfig)
			if err != nil && client != nil {
				log.Fatal("db run err", err)
			}
		})
	}

	return db.client
}

func NewClient(config conf.DBConf) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	var err error
	switch config.Type {
	case conf.TypeMySql:
		db.client, err = OpenMySql(config.Dsn, gormConfig)
	case conf.TypePostgreSQL:
		db.client, err = OpenPostgreSQL(config.Dsn, gormConfig)
	case conf.TypeMSSQL:
		db.client, err = OpenSqlServer(config.Dsn, gormConfig)
	}
	return db.client, err
}
