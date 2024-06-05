package event

// RelationUpdateEvent 关系更新事件
type RelationUpdateEvent struct {
	UserId    int64 `json:"user_id"`
	ThirdId   int64 `json:"third_id"`
	CommentId int64 `json:"comment_id"`
}
