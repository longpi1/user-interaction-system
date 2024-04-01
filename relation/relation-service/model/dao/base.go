package dao

import (
	"sync"

	_ "relation-service/model/dao/db"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var dao *Dao
var once sync.Once

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

func GetDao() *Dao {
	if dao == nil {
		once.Do(func() {
			dao = NewDao()
		})
	}
	return dao
}

// NewDao Create new dao instance This function will auto migrate database tables
func NewDao() *Dao {
	// 启动db与redis
	//dao.db = db.GetClient()
	//dao.cache, _ = cache.Gett
	return dao
}

func (dao *Dao) DB() *gorm.DB {
	return dao.db
}
