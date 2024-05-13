package db

import (
	"sync"

	"github.com/longpi1/user-interaction-system/comment-service/libary/conf"

	"github.com/longpi1/gopkg/libary/log"

	"gorm.io/gorm"
)

var db = &DB{}
var once = sync.Once{}

type DB struct {
	client *gorm.DB
}

func GetClient() *gorm.DB {
	if db == nil || db.client == nil {
		once.Do(func() {
			config := conf.GetConfig()
			client, err := NewClient(config)
			if err != nil && client != nil {
				log.Fatal("db run err", err)
			}
		})
	}

	return db.client
}

func NewClient(config *conf.Config) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	var err error
	switch config.DBConfig.Type {
	case conf.TypeMySql:
		db.client, err = OpenMySql(config.DBConfig.Dsn, gormConfig)
	case conf.TypePostgreSQL:
		db.client, err = OpenPostgreSQL(config.DBConfig.Dsn, gormConfig)
	case conf.TypeMSSQL:
		db.client, err = OpenSqlServer(config.DBConfig.Dsn, gormConfig)
	}
	return db.client, err
}
