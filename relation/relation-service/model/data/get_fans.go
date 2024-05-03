package data

import (
	"fmt"
	"relation-service/model/dao/cache"
	"relation-service/model/dao/db/model"

	"github.com/longpi1/gopkg/libary/log"
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
	relationList, err := model.GetFansList(params)
	if err != nil {
		log.Error("数据库获取粉丝列表败： %v", err)
		return fansList, fmt.Errorf("数据库获取粉丝列表败")
	}
	fansList = formatFansListResponse(relationList)

	// 更新缓存
	cache.SetFansListToLocalCache(fansListLey, fansList)
	cache.SetFansListToRedisCache(fansListLey, fansList)
	return
}

func formatFansListResponse(relations []model.Relation) (fansList model.RelationFansListResponse) {
	fansList.FansCount = len(relations)
	for _, relation := range relations {
		relationResponse := model.RelationResponse{
			UID:        relation.UID,
			Source:     relation.Source,
			Platform:   relation.Platform,
			Status:     relation.Status,
			Ext:        relation.Ext,
			Type:       relation.Type,
			ResourceID: relation.ResourceID,
		}
		fansList.RelationResponse = append(fansList.RelationResponse, relationResponse)
	}
	return fansList
}
