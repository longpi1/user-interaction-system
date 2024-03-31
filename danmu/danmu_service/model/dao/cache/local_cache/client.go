package localcache

import (
	"sync"
	"time"

	"comment-service/libary/conf"
	"comment-service/libary/log"

	"github.com/allegro/bigcache"
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
