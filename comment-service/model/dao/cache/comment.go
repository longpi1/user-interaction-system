package cache

import (
	"fmt"
	"time"
	"user-interaction-system/libary/log"
	localcache "user-interaction-system/model/dao/cache/local_cache"
	"user-interaction-system/model/dao/cache/redis"
	"user-interaction-system/model/dao/db/model"
)

func GetCommentListFromLocalCache(param model.CommentParamsList) (response model.CommentListResponse, err error) {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%s_%s", param.ResourceId, param.Pid, param.Type, param.UserID, param.Content)
	if err = localcache.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetCommentListToLocalCache 将评论列表存入localcache
func SetCommentListToLocalCache(param model.CommentParamsList, response model.CommentListResponse) {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%s_%s", param.ResourceId, param.Pid, param.Type, param.UserID, param.Content)
	err := localcache.Set(key, response, time.Minute*5)
	if err != nil {
		log.Error("评论存入localcache失败: ", key)
	}
}

func GetCommentListFromRedisCache(param model.CommentParamsList) (response model.CommentListResponse, err error) {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%s_%s", param.ResourceId, param.Pid, param.Type, param.UserID, param.Content)
	if err = redis.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetCommentListToRedisCache 将评论列表存入redis
func SetCommentListToRedisCache(param model.CommentParamsList, response model.CommentListResponse) {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%s_%s", param.ResourceId, param.Pid, param.Type, param.UserID, param.Content)
	err := redis.Set(key, response, time.Minute*5)
	if err != nil {
		log.Error("评论存入redis 失败: ", key)
	}
}

// DeleteCommentListCache 删除相关评论列表缓存
func DeleteCommentListCache(param model.CommentParamsAdd) {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%d_%s", param.ResourceId, param.Pid, param.Type, param.UserID, param.Content)
	err := localcache.Delete(key)
	if err != nil {
		log.Error("删除localcache缓存失败", key)
	}
	err = redis.Del(key)
	if err != nil {
		log.Error("删除redis缓存失败", key)
	}
}
