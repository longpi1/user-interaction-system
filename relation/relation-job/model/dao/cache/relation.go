package cache

import (
	"time"

	"github.com/longpi1/user-interaction-system/relation-service/model/dao/cache/redis"
	"github.com/longpi1/user-interaction-system/relation-service/model/dao/db/model"

	"github.com/longpi1/gopkg/libary/log"
)

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
	err := redis.Del(key)
	if err != nil {
		log.Error("删除redis缓存失败", key)
	}
}
