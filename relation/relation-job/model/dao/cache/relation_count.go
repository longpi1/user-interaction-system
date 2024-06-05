package cache

import (
	"time"

	"github.com/longpi1/user-interaction-system/relation-job/model/dao/cache/redis"
	"github.com/longpi1/user-interaction-system/relation-service/model/dao/db/model"

	"github.com/longpi1/gopkg/libary/log"
)

func GetRelationCountFromRedisCache(key string) (response model.RelationCountResponse, err error) {
	if err = redis.Get(key, response); err != nil {
		log.Error("获取redis评论数失败: ", key)
		return response, err
	}
	return response, nil
}

// SetRelationCountToRedisCache 将关注数信息存入redis
func SetRelationCountToRedisCache(key string, response model.RelationCountResponse) {
	err := redis.Set(key, response, time.Minute*120)
	if err != nil {
		log.Error("关注数存入redis 失败: ", key)
	}
}

// DeleteRelationCountCache 删除相关关注数缓存
func DeleteRelationCountCache(key string) {
	err := redis.Del(key)
	if err != nil {
		log.Error("删除redis缓存失败", key)
	}
}
