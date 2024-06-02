package msg

// 消息状态
const (
	StatusUnread   = 0 // 消息未读
	StatusHaveRead = 1 // 消息已读
)

type Type int

// 消息类型
const (
	TypeResourceComment   Type = 0 // 收到资源评论
	TypeCommentReply      Type = 1 // 收到他人回复
	TypeResourceLike      Type = 2 // 收到点赞
	TypeResourceFavorite  Type = 3 // 话题被收藏
	TypeResourceRecommend Type = 4 // 话题被设为推荐
	TypeResourceDelete    Type = 5 // 话题被删除
)

type ResourceLikeExtraData struct {
	ResourceId int64 `json:"resource_id"`
	LikeUserId int64 `json:"likeUserId"`
}

type ResourceFavoriteExtraData struct {
	ResourceId     int64 `json:"resource_id"`
	FavoriteUserId int64 `json:"favoriteUserId"`
}

type ResourceRecommendExtraData struct {
	ResourceId int64 `json:"resource_id"`
}

type ResourceDeleteExtraData struct {
	ResourceId   int64 `json:"resource_id"`
	DeleteUserId int64 `json:"deleteUserId"`
}

type CommentExtraData struct {
	ResourceType string `json:"resourceType"` // 评论实体类型
	ResourceID   int64  `json:"resource_id"`  // 评论实体ID
	PID          int64  `json:"pid"`          // 引用评论ID
	RootType     string `json:"rootType"`     // 根评论的实体类型（例如：文章评论的二级评论，该类型还是文章）
	RootId       string `json:"root_id"`      // 根评论的实体ID
}
