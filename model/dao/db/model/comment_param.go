package model

type CommentParamsList struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	Username string `json:"username"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
	OrderBy  string `json:"oderby"`
}

type CommenParamsAdd struct {
	UserName      string `form:"username"`
	Content       string `form:"content" validate:"required"`
	ResourceId    uint   `form:"content" validate:"required"` // 评论所关联的资源id
	ResourceTitle string `form:"content" validate:"required"` // 资源的title
	ContentMeta   string `form:"cotent_meta"`                 // 存储一些关键的附属信息
	Pid           uint   `form:"pid"`                         // 父评论 ID
	UserID        uint   `form:"user_id" validate:"required"` //  发表者id
	Ext           string `form:"ext"`
	IP            string `form:"ip"`
	ContentRich   string `form:"content_rich"`
	Type          uint   `form:"type"` // 评论类型，文字评论、图评等"`                         // 额外信息存储
}
