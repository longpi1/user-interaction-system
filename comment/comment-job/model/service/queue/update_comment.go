package queue

import "comment-job/libary/event"

func UpdateComment() {
	// 发送事件
	event.Send(event.CommentUpdateEvent{
		UserId:    userId,
		CommentId: comment.Id,
	})
}
