package queue

import (
	"github.com/longpi1/user-interaction-system/comment-job/libary/event"
	"github.com/longpi1/user-interaction-system/comment-job/model/dao/cache"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"
)

func UpdateComment(commentInfo model.CommentInfo) error {
	// todo redis相关更新操作
	// 删除相关评论列表缓存数据
	key := cache.GetCommentListKey(commentInfo.ResourceId, commentInfo.Pid)
	cache.DeleteCommentListCache(key)
	// 更新
	// 发送事件，进行后置更新，消息通知等行为
	event.Send(event.CommentUpdateEvent{
		UserId:     commentInfo.UserID,
		ResourceId: commentInfo.ResourceId,
		CommentId:  commentInfo.CommentId,
	})
	return nil
}
