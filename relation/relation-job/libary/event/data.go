package event

// CommentUpdateEvent 评论更新事件
type CommentUpdateEvent struct {
	UserId     int64 `json:"user_id"`
	ResourceId int64 `json:"resource_id"`
	CommentId  int64 `json:"comment_id"`
}
