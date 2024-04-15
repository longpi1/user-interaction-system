package data

import (
	"fmt"
	"relation-service/libary/log"
	"relation-service/model/dao/cache"
	"relation-service/model/dao/db/model"
)

func GetFollowingList(params model.RelationFollowingParams) {
	followingListKey := cache.GetRelationFollowingListKey(params.UID, relation.Platform, relation.Type, params.Status)
	// 首先从缓存中获取对应数据
	if countResponse, err = cache.Get(key); err == nil {
		return countResponse, nil
	}
	if countResponse, err = cache.GetRelationCountFromRedisCache(key); err == nil {
		return countResponse, nil
	}
	// 缓存中获取失败则从数据库中获取
	relationCount, err := model.FindRelationCountByParams(params)
	if err != nil {
		log.Error("数据库获取关注数、粉丝数失败： %v", err)
		return countResponse, fmt.Errorf("数据库获取关注数、粉丝数失败")
	}
	countResponse = formatRelationCountResponse(relationCount)

	// 更新缓存
	cache.SetRelationCountToLocalCache(key, countResponse)
	cache.SetRelationCountToRedisCache(key, countResponse)
	return
}
