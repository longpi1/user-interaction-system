package model

type CommentInfo struct {
	CommentId     int64  // 评论id
	ResourceId    int64  // 评论所关联的资源id
	ResourceTitle string // 资源的title
	Content       string // 文本信息
	ContentMeta   string // 存储一些关键的附属信息
	ContentRich   string
	Pid           int64  // 父评论 ID
	Ext           string // 额外信息存储
	IP            string
	IPArea        string
	UserID        int64  // 发表者id
	UserName      string // 发表者名称
}
