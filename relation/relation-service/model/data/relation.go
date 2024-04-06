package data

import (
	"fmt"
	"relation-service/libary/conf"
	"relation-service/libary/log"
	"relation-service/model/dao/cache"
	"relation-service/model/dao/db"
	"relation-service/model/dao/db/model"
)

func Follow(params model.RelationParams) error {
	relation := formatRelation(params)

	tx := db.GetClient().Begin()
	_, err := model.InsertRelationWithTx(tx, &relation)
	if err != nil {
		log.Error("数据库插入失败", err)
		return fmt.Errorf("数据库插入失败")
	}
	// 删除原有缓存
	key := cache.GetRelationListKey(params.UID, params.Platform, params.Type)
	cache.DeleteRelationCache(key)
	return nil
}

func UnFollow(params model.RelationParams) error {
	tx := db.GetClient().Begin()
	err := model.DeleteRelationWithTx(tx, params.UID, params.ResourceID)
	if err != nil {
		log.Error("数据库删除失败", err)
		return fmt.Errorf("数据库删除失败")
	}
	// 删除原有缓存
	key := cache.GetRelationListKey(params.UID, params.Platform, params.Type)
	cache.DeleteRelationCache(key)
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
