package data

import (
	"fmt"
	"relation-service/libary/constant"
	"relation-service/model/dao/cache"
	"relation-service/model/dao/db/model"

	"github.com/longpi1/gopkg/libary/utils"

	"github.com/longpi1/gopkg/libary/log"
)

func IsFollowing(params model.RelationIsFollowingParams) (model.RelationIsFollowingResponse, error) {
	platform := utils.ConvertPlatform(params.Platform)
	relationType := utils.ConvertPlatform(params.Type)
	followingListLey := cache.GetFollowingListKey(params.UID, platform, relationType, constant.FollowingStatus)
	// 首先从缓存中获取对应数据
	if followingList, err := cache.GetFollowingListFromLocalCache(followingListLey); err == nil {
		return judgeFollowing(params.ResourceID, followingList), nil
	}
	if followingList, err := cache.GetFollowingListFromRedisCache(followingListLey); err == nil {
		return judgeFollowing(params.ResourceID, followingList), nil
	}
	// 缓存中获取失败则从数据库中获取
	relation, err := model.GetIsFollowing(params.UID, relationType, platform, params.ResourceID)
	if err != nil {
		log.Error("数据库获取关注列表失败： %v, %d", err, params.ResourceID)
		return model.RelationIsFollowingResponse{}, fmt.Errorf("数据库获取关注关系失败")
	}
	if relation == nil {
		return model.RelationIsFollowingResponse{}, nil
	}
	return model.RelationIsFollowingResponse{IsFollowing: true}, nil
}

func IsFollowingBatch(params model.RelationIsFollowingBatchParams) (model.RelationIsFollowingBatchResponse, error) {
	var batchResponse model.RelationIsFollowingBatchResponse
	platform := utils.ConvertPlatform(params.Platform)
	relationType := utils.ConvertPlatform(params.Type)
	followingListLey := cache.GetFollowingListKey(params.UID, platform, relationType, constant.FollowingStatus)
	// 首先从缓存中获取对应数据
	if followingList, err := cache.GetFollowingListFromLocalCache(followingListLey); err == nil {
		for _, resourceId := range params.ResourceIDs {
			batchResponse.Data[params.UID] = judgeFollowing(resourceId, followingList)
		}
	}
	if followingList, err := cache.GetFollowingListFromRedisCache(followingListLey); err == nil {
		for _, resourceId := range params.ResourceIDs {
			batchResponse.Data[params.UID] = judgeFollowing(resourceId, followingList)
		}
	}
	for _, resourceId := range params.ResourceIDs {
		// 缓存中获取失败则从数据库中获取
		relation, err := model.GetIsFollowing(params.UID, relationType, platform, resourceId)
		if err != nil {
			log.Error("数据库获取关注列表失败： %v, %d", err, resourceId)
			continue
		}
		if relation == nil {
			batchResponse.Data[params.UID] = model.RelationIsFollowingResponse{}
			continue
		}
		batchResponse.Data[params.UID] = model.RelationIsFollowingResponse{IsFollowing: true}
	}
	return batchResponse, nil
}

func judgeFollowing(resourceID int64, response model.RelationFollowingListResponse) model.RelationIsFollowingResponse {
	for _, relationResponse := range response.RelationResponse {
		if resourceID == relationResponse.ResourceID {
			return model.RelationIsFollowingResponse{
				IsFollowing: true,
			}
		}
	}
	return model.RelationIsFollowingResponse{
		IsFollowing: false,
	}
}
