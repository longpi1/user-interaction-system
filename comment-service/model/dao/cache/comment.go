package cache

import (
	"fmt"
	"time"
	localcache "user-interaction-system/model/dao/cache/local_cache"
	"user-interaction-system/model/dao/db/model"
)

// 从localcache中获取评论列表
var commentLocalCache = localcache.NewCache[model.CommentListResponse]()

func GetCommentListFromLocalCache(param model.CommentParamsList) (model.CommentListResponse, error) {
	cache := localcache.GetClient()
	key := fmt.Sprintf("comment_list_%d_%s_%d_%d_%s_%s", param.ResourceId, param.ResourceTitle, param.Pid, param.Type, param.Username, param.Content)
	if response, ok := cache.Get(key); ok {
		return response, nil
	}
	return model.CommentListResponse{}, nil
}

// 将评论列表存入localcache
func SetCommentListToLocalCache(key string, response model.CommentListResponse) {
	commentLocalCache.Set(key, response, time.Minute*5)
}

// 从redis中获取评论列表
var commentRedisCache = rediscache.NewCache[model.CommentListResponse]()

func GetCommentListFromRedisCache(param model.CommentParamsList) (model.CommentListResponse, error) {
	key := fmt.Sprintf("comment_list_%d_%s_%d_%d_%s_%s", param.ResourceId, param.ResourceTitle, param.Pid, param.Type, param.Username, param.Content)
	if response, ok := commentRedisCache.Get(key); ok {
		return response, nil
	}
	return model.CommentListResponse{}, nil
}

// 将评论列表存入redis
func SetCommentListToRedisCache(key string, response model.CommentListResponse) {
	commentRedisCache.Set(key, response, time.Minute*5)
}
