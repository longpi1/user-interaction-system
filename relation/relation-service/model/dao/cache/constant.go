package cache

import "fmt"

const (
	RelationListRedisKey = "relation_list_%d_%s_%s"
)

func GetRelationListKey(UID int64, platform string, relationType string) string {
	key := fmt.Sprintf(RelationListRedisKey, UID, platform, relationType)
	return key
}
