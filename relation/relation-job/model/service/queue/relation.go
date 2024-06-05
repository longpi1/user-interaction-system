package queue

import (
	"github.com/longpi1/user-interaction-system/relation-job/libary/event"
	"github.com/longpi1/user-interaction-system/relation-job/model/dao/cache"
	"github.com/longpi1/user-interaction-system/relation-service/model/dao/db/model"
)

func Relation(relationInfo model.Relation) error {
	// todo 关注数相关更新操作

	// 删除相关关注列表缓存数据
	followingListKey := cache.GetFollowingListKey(relationInfo.UID, relationInfo.Platform, relationInfo.Type, relationInfo.Status)
	fansListLey := cache.GetFansListKey(relationInfo.ResourceID, relationInfo.Platform, relationInfo.Type, relationInfo.Status)
	cache.DeleteRelationCache(followingListKey)
	cache.DeleteRelationCache(fansListLey)
	// 发送事件，进行后置更新，消息通知等行为
	event.Send(event.RelationUpdateEvent{})
	return nil
}
