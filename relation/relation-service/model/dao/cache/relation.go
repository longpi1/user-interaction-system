package cache

import (
	"relation-service/libary/log"
	localcache "relation-service/model/dao/cache/local_cache"
	"relation-service/model/dao/cache/redis"
	"relation-service/model/dao/db/model"
	"time"
)

func GetRelationFromLocalCache(key string) (response model.RelationListResponse, err error) {
	if err = localcache.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetRelationToLocalCache 将评论列表存入localcache
func SetRelationToLocalCache(key string, response model.RelationListResponse) {
	err := localcache.Set(key, response, time.Minute*5)
	if err != nil {
		log.Error("评论存入localcache失败: ", key)
	}
}

func GetRelationFromRedisCache(key string) (response model.RelationListResponse, err error) {
	if err = redis.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetRelationToRedisCache 将关注列表存入redis
func SetRelationToRedisCache(key string, response model.RelationListResponse) {
	err := redis.Set(key, response, time.Minute*5)
	if err != nil {
		log.Error("关注存入redis 失败: ", key)
	}
}

// DeleteRelationCache 删除相关关注列表缓存
func DeleteRelationCache(key string) {
	err := localcache.Delete(key)
	if err != nil {
		log.Error("删除localcache缓存失败", key)
	}
	err = redis.Del(key)
	if err != nil {
		log.Error("删除redis缓存失败", key)
	}
}
