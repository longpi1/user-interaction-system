package cache

import "fmt"

const (
	// RelationFollowingListRedisKey 关注列表缓存
	RelationFollowingListRedisKey = "relation_follow_list_%d_%d_%d_%d"
	// RelationFansListRedisKey 粉丝列表缓存
	RelationFansListRedisKey = "relation_fan_list_%d_%d_%d_%d"
	//RelationCountRedisKey 关注数缓存
	RelationCountRedisKey = "relation_count_%d_%d_%d"
)

func GetFollowingListKey(UID int64, platform int, relationType int, status int) string {
	// -1 表示获取所有状态
	if status == -1 {
		return fmt.Sprintf(RelationFollowingListRedisKey, UID, platform, relationType, "*")
	}
	key := fmt.Sprintf(RelationFollowingListRedisKey, UID, platform, relationType, status)
	return key
}

func GetFansListKey(ResourceId int64, platform int, relationType int, status int) string {
	key := fmt.Sprintf(RelationFansListRedisKey, ResourceId, platform, relationType, status)
	return key
}

func GetRelationCountKey(ResourceID int64, platform int, relationType int) string {
	key := fmt.Sprintf(RelationCountRedisKey, ResourceID, platform, relationType)
	return key
}
