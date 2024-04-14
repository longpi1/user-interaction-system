package cache

import "fmt"

const (
	// RelationListRedisKey 关注列表缓存
	RelationListRedisKey = "relation_list_%d_%d_%d_%d"
	//RelationCountRedisKey 关注数缓存
	RelationCountRedisKey = "relation_count_%d_%d_%d"
)

func GetRelationListKey(UID int64, platform int, relationType int, status int) string {
	key := fmt.Sprintf(RelationListRedisKey, UID, platform, relationType, status)
	return key
}

func GetRelationCountKey(ResourceID int64, platform int, relationType int) string {
	key := fmt.Sprintf(RelationCountRedisKey, ResourceID, platform, relationType)
	return key
}
