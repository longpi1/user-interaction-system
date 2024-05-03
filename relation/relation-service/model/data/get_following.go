package data

import (
	"fmt"
	"relation-service/model/dao/cache"
	"relation-service/model/dao/db/model"

	"github.com/longpi1/gopkg/libary/log"
)

func GetFollowingList(params model.RelationFollowingParams) (followingList model.RelationFollowingListResponse, err error) {
	followingListLey := cache.GetFollowingListKey(params.UID, params.Platform, params.Type, params.Status)
	// 首先从缓存中获取对应数据
	if followingList, err = cache.GetFollowingListFromLocalCache(followingListLey); err == nil {
		return followingList, nil
	}
	if followingList, err = cache.GetFollowingListFromRedisCache(followingListLey); err == nil {
		return followingList, nil
	}
	// 缓存中获取失败则从数据库中获取
	relationList, err := model.GetFollowingList(params)
	if err != nil {
		log.Error("数据库获取关注列表失败： %v", err)
		return followingList, fmt.Errorf("数据库获取关注列表失败")
	}
	followingList = formatFollowingListResponse(relationList)

	// 更新缓存
	cache.SetFollowingListToLocalCache(followingListLey, followingList)
	cache.SetFollowingListToRedisCache(followingListLey, followingList)
	return
}

func formatFollowingListResponse(relations []model.Relation) (followingList model.RelationFollowingListResponse) {
	followingList.FollowingCount = len(relations)
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
		followingList.RelationResponse = append(followingList.RelationResponse, relationResponse)
	}
	return followingList
}
