package cache

import (
	"time"

	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"

github.com/longpi1/user-interaction-system/comment-job"/model/dao/cache/redis"

"github.com/longpi1/gopkg/libary/log"
)

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
	err := redis.Del(key)
	if err != nil {
		log.Error("删除redis缓存失败", key)
	}
}
