package queue

import (
	"github.com/longpi1/user-interaction-system/comment/comment-job/job/comment"
	"github.com/longpi1/user-interaction-system/comment/comment-job/libary/event"
)

func UpdateComment(commentInfo comment.CommentInfo) error {
	// todo redis相关更新操作

	// 发送事件，进行后置更新，消息通知等行为
	event.Send(event.CommentUpdateEvent{
		UserId:     commentInfo.UserID,
		ResourceId: commentInfo.ResourceId,
		CommentId:  commentInfo.CommentId,
	})
	return nil
}
