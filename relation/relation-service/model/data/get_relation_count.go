package data

import (
	"fmt"
	"relation-service/libary/log"
	"relation-service/model/dao/cache"
	"relation-service/model/dao/db/model"
)

func GetRelationCount(params model.RelationCountParams) (countResponse model.RelationCountResponse, err error) {
	key := cache.GetRelationCountKey(params.ResourceID, params.Platform, params.Type)
	// 首先从缓存中获取对应数据
	if countResponse, err = cache.GetRelationCountFromLocalCache(key); err == nil {
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

func formatRelationCountResponse(relationCount model.RelationCount) model.RelationCountResponse {
	return model.RelationCountResponse{
		ResourceID:  relationCount.ResourceId,
		Platform:    int(relationCount.Platform),
		Type:        int(relationCount.Type),
		FollowCount: int(relationCount.FollowCount),
		FansCount:   int(relationCount.FansCount),
	}
}
