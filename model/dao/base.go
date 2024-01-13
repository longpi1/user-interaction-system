package dao

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"model-api/libary/conf"
	"model-api/libary/log"
	"model-api/model/dao/db"
)

type Dao struct {
	db *gorm.DB

	// Cache to speed up database query
	//
	// Please access the cache instance by invoking the `CacheAction` func.
	//
	// As the cache could be nil.
	// Therefore, it is necessary to check if the cache is not nil before referencing it.
	cache *redis.Client
}

// Create new dao instance
//
// This function will auto migrate database tables
func NewDao(config conf.DBConf) *Dao {
	// 启动db与redis
	db,err := db.NewClient(config)
	if err != nil {
		log.Fatal("db run err", err)
	}
	dao := &Dao{
		db: db,
	}

	return dao
}

func (dao *Dao) DB() *gorm.DB {
	return dao.db
}

