package cache

import "fmt"

func GetCommentListKey(resourceId int, pid int, sourceType int, userId int) string {
	key := fmt.Sprintf("comment_list_%d_%d_%d_%d", resourceId, pid, sourceType, userId)
	return key
}
