package cache

import (
	"time"

	localcache "github.com/longpi1/user-interaction-system/comment-service/model/dao/cache/local_cache"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/cache/redis"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"

	"github.com/longpi1/gopkg/libary/log"
)

func GetCommentListFromLocalCache(key string) (response model.CommentListResponse, err error) {
	if err = localcache.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetCommentListToLocalCache 将评论列表存入localcache
func SetCommentListToLocalCache(key string, response model.CommentListResponse) {
	err := localcache.Set(key, response, time.Minute*5)
	if err != nil {
		log.Error("评论存入localcache失败: ", key)
	}
}

func GetCommentListFromRedisCache(key string) (response model.CommentListResponse, err error) {
	if err = redis.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetCommentListToRedisCache 将评论列表存入redis
func SetCommentListToRedisCache(key string, response model.CommentListResponse) {
	err := redis.Set(key, response, time.Minute*5)
	if err != nil {
		log.Error("评论存入redis 失败: ", key)
	}
}

// DeleteCommentListCache 删除相关评论列表缓存
func DeleteCommentListCache(key string) {
	err := localcache.Delete(key)
	if err != nil {
		log.Error("删除localcache缓存失败", key)
	}
	err = redis.Del(key)
	if err != nil {
		log.Error("删除redis缓存失败", key)
	}
}
