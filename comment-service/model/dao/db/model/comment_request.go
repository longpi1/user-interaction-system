package model

type CommentParamsList struct {
	ResourceId uint   `form:"content" validate:"required"` // 评论所关联的资源id
	Pid        uint   `form:"pid"`                         // 父评论 ID
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	UserID     uint   `form:"user_id" validate:"required"`
	Content    string `json:"content"`
	Type       int    `json:"type"`
	OrderBy    string `json:"oderby"`
}

type CommentParamsAdd struct {
	UserName      string `form:"username"`
	Content       string `form:"content" validate:"required"`
	ResourceId    uint   `form:"resource_id" validate:"required"`    // 评论所关联的资源id
	ResourceTitle string `form:"resource_title" validate:"required"` // 资源的title
	ContentMeta   string `form:"content_meta"`                       // 存储一些关键的附属信息
	Pid           uint   `form:"pid"`                                // 父评论 ID
	UserID        uint   `form:"user_id" validate:"required"`        //  发表者id
	Ext           string `form:"ext"`                                // 额外信息存储
	IP            string `form:"ip"`
	ContentRich   string `form:"content_rich"`
	Type          uint   `form:"type"` // 评论类型，文字评论、图评等"`
}
