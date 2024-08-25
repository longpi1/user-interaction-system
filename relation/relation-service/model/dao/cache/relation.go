package cache

import (
	localcache "relation-service/model/dao/cache/local_cache"
	"relation-service/model/dao/cache/redis"
	"relation-service/model/dao/db/model"
	"time"

	"github.com/longpi1/gopkg/libary/log"
)

func GetFansListFromLocalCache(key string) (response model.RelationFansListResponse, err error) {
	if err = localcache.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetFansListToLocalCache 将粉丝列表存入localcache
func SetFansListToLocalCache(key string, response model.RelationFansListResponse) {
	err := localcache.Set(key, response, time.Minute*120)
	if err != nil {
		log.Error("粉丝列表存入localcache失败: ", key)
	}
}

func GetFansListFromRedisCache(key string) (response model.RelationFansListResponse, err error) {
	if err = redis.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetFansListToRedisCache 将关注列表存入redis
func SetFansListToRedisCache(key string, response model.RelationFansListResponse) {
	err := redis.Set(key, response, time.Minute*120)
	if err != nil {
		log.Error("粉丝列表存入redis 失败: ", key)
	}
}

func GetFollowingListFromLocalCache(key string) (response model.RelationFollowingListResponse, err error) {
	if err = localcache.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetFollowingListToLocalCache 将关注列表存入localcache
func SetFollowingListToLocalCache(key string, response model.RelationFollowingListResponse) {
	err := localcache.Set(key, response, time.Minute*120)
	if err != nil {
		log.Error("关注列表存入localcache失败: ", key)
	}
}

func GetFollowingListFromRedisCache(key string) (response model.RelationFollowingListResponse, err error) {
	if err = redis.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetFollowingListToRedisCache 将关注列表存入redis
func SetFollowingListToRedisCache(key string, response model.RelationFollowingListResponse) {
	err := redis.Set(key, response, time.Minute*120)
	if err != nil {
		log.Error("关注列表存入redis 失败: ", key)
	}
}

// DeleteRelationCache 删除相关关注列表缓存
func DeleteRelationCache(key string) {
	err := localcache.Delete(key)
	if err != nil {
		log.Error("删除localcache缓存失败", key)
	}
}
