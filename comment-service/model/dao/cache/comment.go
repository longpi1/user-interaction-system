package cache

import (
	"comment-service/libary/log"
	localcache "comment-service/model/dao/cache/local_cache"
	"comment-service/model/dao/cache/redis"
	"comment-service/model/dao/db/model"
	"fmt"
	"time"
)

func GetCommentListFromLocalCache(param model.CommentParamsList) (response model.CommentListResponse, err error) {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%d", param.ResourceId, param.Pid, param.Type, param.UserID)
	if err = localcache.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetCommentListToLocalCache 将评论列表存入localcache
func SetCommentListToLocalCache(param model.CommentParamsList, response model.CommentListResponse) {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%d", param.ResourceId, param.Pid, param.Type, param.UserID)
	err := localcache.Set(key, response, time.Minute*5)
	if err != nil {
		log.Error("评论存入localcache失败: ", key)
	}
}

func GetCommentListFromRedisCache(param model.CommentParamsList) (response model.CommentListResponse, err error) {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%d", param.ResourceId, param.Pid, param.Type, param.UserID)
	if err = redis.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetCommentListToRedisCache 将评论列表存入redis
func SetCommentListToRedisCache(param model.CommentParamsList, response model.CommentListResponse) {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%d", param.ResourceId, param.Pid, param.Type, param.UserID)
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
