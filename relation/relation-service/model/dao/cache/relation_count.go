package cache

import (
	"relation-service/libary/log"
	localcache "relation-service/model/dao/cache/local_cache"
	"relation-service/model/dao/cache/redis"
	"relation-service/model/dao/db/model"
	"time"
)

func GetRelationCountFromLocalCache(key string) (response model.RelationCountResponse, err error) {
	if err = localcache.Get(key, response); err != nil {
		log.Error("获取localcache评论数失败: ", key)
		return response, err
	}
	return response, nil
}

// SetRelationCountToLocalCache 将关注数信息存入本地缓存
func SetRelationCountToLocalCache(key string, response model.RelationCountResponse) {
	err := localcache.Set(key, response, time.Minute*120)
	if err != nil {
		log.Error("评论数存入localcache失败: ", key)
	}
}

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
	err := localcache.Delete(key)
	if err != nil {
		log.Error("删除localcache缓存失败", key)
	}
	err = redis.Del(key)
	if err != nil {
		log.Error("删除redis缓存失败", key)
	}
}
