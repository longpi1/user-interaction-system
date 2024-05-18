package model

type CommentInfo struct {
	CommentId     int64  // 评论id
	ResourceId    int64  // 评论所关联的资源id
	ResourceType  string // 资源的title
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
	Type          int    // 消息类型，0新增、1删除、2审核通过，3审核拒绝、4修改等
}

func FormatCommentInfo(commentIndex CommentIndex, commentContent CommentContent, opType int) CommentInfo {
	return CommentInfo{
		CommentId:    int64(commentIndex.ID),
		ResourceId:   commentIndex.ResourceID,
		ResourceType: commentIndex.ResourceType,
		Content:      commentContent.Content,
		ContentMeta:  commentContent.ContentMeta,
		ContentRich:  commentContent.ContentRich,
		Pid:          commentIndex.PID,
		Ext:          commentContent.Ext,
		UserID:       commentIndex.UserID,
		UserName:     commentIndex.UserName,
		IP:           commentIndex.IP,
		IPArea:       commentIndex.IPArea,
		Type:         opType,
	}
}
