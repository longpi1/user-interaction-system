package localcache

import (
	"github.com/allegro/bigcache"
	"sync"
	"time"
	"user-interaction-system/libary/conf"
	"user-interaction-system/libary/log"
)

var localCache *bigcache.BigCache
var once sync.Once

func GetClient() *bigcache.BigCache {
	var err error
	if localCache == nil {
		once.Do(func() {
			localCache, err = NewClient(conf.GetConfig())
			if err != nil {
				log.Fatal("local cache run err", err)
			}
		})
	}
	return localCache
}

func NewClient(config *conf.Config) (*bigcache.BigCache, error) {
	var err error
	evictionTime := config.LocalCache.EvictionTime
	localCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(evictionTime * time.Minute))
	return localCache, err
}
