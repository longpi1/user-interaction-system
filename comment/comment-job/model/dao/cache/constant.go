package cache

import "fmt"

const (
	CommentListRedisKey           = "comment_list_%d_%d"
	CommentListRedisKeyWithNilPid = "comment_list_%d_*"
)

func GetCommentListKey(resourceId int64, pid int64) string {
	// 如果pid为默认值也就是零则用_*
	if pid == 0 {
		key := fmt.Sprintf(CommentListRedisKeyWithNilPid, resourceId)
		return key
	}
	key := fmt.Sprintf(CommentListRedisKey, resourceId, pid)
	return key
}
