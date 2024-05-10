package model

type CommentResponse struct {
	CommentId     uint   // 评论id
	ResourceId    uint   // 评论所关联的资源id
	ResourceTitle string // 资源的title
	Content       string // 文本信息
	ContentMeta   string // 存储一些关键的附属信息
	ContentRich   string
	Pid           uint   // 父评论 ID
	Ext           string // 额外信息存储
	IP            string
	IPArea        string
	UserID        uint   // 发表者id
	UserName      string // 发表者名称
	Type          uint   // 评论类型，文字评论、图评等
	IsCollapsed   bool   // 折叠
	IsPending     bool   // 待审
	IsPinned      bool   // 置顶
	State         uint   // 状态
	LikeCount     uint   // 点赞数
	HateCount     uint   // 点踩数
	ReplyCount    uint   // 回复数

	FloorCount uint // 评论层数
}

type CommentListResponse struct {
	CommentResponses []CommentResponse
	RootReplyCount   uint // 根评论回复数
}
