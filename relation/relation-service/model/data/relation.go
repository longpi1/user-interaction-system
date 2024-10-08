package data

import (
	"fmt"
	"relation-service/libary/conf"
	"relation-service/model/dao/cache"
	"relation-service/model/dao/db"
	"relation-service/model/dao/db/model"

	"github.com/longpi1/gopkg/libary/log"
	"github.com/longpi1/gopkg/libary/queue"
	"github.com/longpi1/gopkg/libary/utils"
)

func Follow(params model.RelationParams) error {
	relation := formatRelation(params)
	tx := db.GetClient().Begin()
	_, err := model.InsertRelationWithTx(tx, &relation)
	if err != nil {
		log.Error("数据库更新失败", err)
		return fmt.Errorf("数据库更新失败")
	}
	// 队列发送关注数据进行后置更新
	queueConfig := conf.GetConfig().QueueConfig
	if err := queue.Push(queueConfig.TopicName, relation, queueConfig.Config); err != nil {
		log.Error("队列发送数据失败：%v", relation)
	}
	// 删除原有缓存
	followingListKey := cache.GetFollowingListKey(params.UID, relation.Platform, relation.Type, params.Status)
	fansListLey := cache.GetFansListKey(params.ResourceID, relation.Platform, relation.Type, params.Status)
	cache.DeleteRelationCache(followingListKey)
	cache.DeleteRelationCache(fansListLey)
	// todo 通过job去更新粉丝数、关注数数据库与缓存，如果分平台则需要更新整体数量
	return nil
}

func UnFollow(params model.RelationParams) error {
	relation := formatRelation(params)
	tx := db.GetClient().Begin()
	err := model.DeleteRelationWithTx(tx, params.UID, params.ResourceID)
	if err != nil {
		log.Error("数据库删除失败", err)
		return fmt.Errorf("数据库删除失败")
	}
	// 队列发送关注数据进行后置更新
	queueConfig := conf.GetConfig().QueueConfig
	if err := queue.Push(queueConfig.TopicName, relation, queueConfig.Config); err != nil {
		log.Error("队列发送数据失败：%v", relation)
	}
	// 删除原有缓存
	followingListKey := cache.GetFollowingListKey(params.UID, utils.ConvertPlatform(params.Platform), utils.ConvertType(params.Type), params.Status)
	fansListLey := cache.GetFansListKey(params.ResourceID, utils.ConvertPlatform(params.Platform), utils.ConvertType(params.Type), params.Status)
	cache.DeleteRelationCache(followingListKey)
	cache.DeleteRelationCache(fansListLey)
	// todo 通过job去更新粉丝数、关注数数据库与缓存，如果分平台则需要更新整体数量
	return nil
}

func formatRelation(params model.RelationParams) model.Relation {
	relation := model.Relation{
		Type:       getRelationType(params.Type),
		Platform:   getRelationPlatform(params.Platform),
		Ext:        params.Ext,
		UID:        params.UID,
		ResourceID: params.ResourceID,
		Source:     params.Source,
	}
	return relation
}

// getRelationType 将参数的string类型改为int类型存入数据库
func getRelationType(relationType string) int {
	configMap := conf.GetMapConfig()
	result := configMap.TypeMap[relationType]
	return result
}

// getRelationPlatform 将参数的string类型改为int类型存入数据库
func getRelationPlatform(relationPlatform string) int {
	configMap := conf.GetMapConfig()
	result := configMap.PlatformMap[relationPlatform]
	return result
}
