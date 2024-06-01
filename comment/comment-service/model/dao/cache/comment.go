package cache

import (
	"math"
	"time"

	"github.com/longpi1/gopkg/libary/log"
	localcache "github.com/longpi1/user-interaction-system/comment-service/model/dao/cache/local_cache"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/cache/redis"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"

	"github.com/go-kratos/aegis/topk"
)

const (
	HotKeyCnt = 10
	MinCount  = 0
)

var resourceTopk topk.Topk

func GetCommentListFromLocalCache(key string) (response model.CommentListResponse, err error) {
	if resourceTopk != nil {
		_, _ = resourceTopk.Add(key, 1)
	}
	if err = localcache.Get(key, response); err != nil {
		return response, err
	}
	return response, nil
}

// SetCommentListToLocalCache 将评论列表存入localcache
func SetCommentListToLocalCache(key string, response model.CommentListResponse) {
	var added bool
	if resourceTopk != nil {
		_, added = resourceTopk.Add(key, 1)
	}
	if added {
		err := localcache.Set(key, response, time.Hour*24)
		if err != nil {
			log.Error("评论存入localcache失败: ", key)
		}
	} else {
		if inWhileList(key) {
			err := localcache.Set(key, response, time.Hour*24)
			if err != nil {
				log.Error("评论存入localcache失败: ", key)
			}
		}
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
	err := redis.Set(key, response, time.Hour*24)
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

func inWhileList(key string) bool {
	return true
}

func init() {
	factor := uint32(math.Log(float64(HotKeyCnt)))
	resourceTopk = topk.NewHeavyKeeper(uint32(HotKeyCnt), 1024*factor, 4, 0.925, uint32(MinCount))
}
