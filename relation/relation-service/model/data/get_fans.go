package data

import (
	"fmt"
	"relation-service/libary/log"
	"relation-service/model/dao/cache"
	"relation-service/model/dao/db/model"
)

func GetFansList(params model.RelationFansParams) (fansList model.RelationFansListResponse, err error) {
	fansListLey := cache.GetFansListKey(params.ResourceID, params.Platform, params.Type, params.Status)
	// 首先从缓存中获取对应数据
	if fansList, err = cache.GetFansListFromLocalCache(fansListLey); err == nil {
		return fansList, nil
	}
	if fansList, err = cache.GetFansListFromRedisCache(fansListLey); err == nil {
		return fansList, nil
	}
	// 缓存中获取失败则从数据库中获取
	relationCount, err := model.Fin(params)
	if err != nil {
		log.Error("数据库获取关注数、粉丝数失败： %v", err)
		return fansList, fmt.Errorf("数据库获取关注数、粉丝数失败")
	}
	countResponse = formatRelationCountResponse(relationCount)

	// 更新缓存
	cache.SetFansListToLocalCache(fansListLey, fansList)
	cache.SetFansListToRedisCache(fansListLey, fansList)
	return
}

func formatFansListResponse() {

}
