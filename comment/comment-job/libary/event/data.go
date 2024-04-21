package event

// CommentCreateEvent 评论创建事件
type CommentCreateEvent struct {
	UserId     int64 `json:"user_id"`
	ResourceId int64 `json:"resource_id"`
	CommentId  int64 `json:"comment_id"`
}
