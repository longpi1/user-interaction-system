package model

type CommentParamsList struct {
	ResourceId uint   `form:"content" validate:"required"` // 评论所关联的资源id
	Pid        uint   `form:"pid"`                         // 父评论 ID
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
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

type CommentParamsDelete struct {
	ResourceId uint   `form:"resource_id" validate:"required"` // 评论所关联的资源id
	CommentID  uint   `form:"comment_id"`                      // 父评论 ID
	UserID     uint   `form:"user_id" validate:"required"`     //  发表者id
	Ext        string `form:"ext"`                             // 额外信息存储
	Pid        uint   `form:"pid"`                             // 父评论 ID
}

type CommentParamsInteract struct {
	ResourceId uint   `form:"resource_id" validate:"required"` // 评论所关联的资源id
	CommentID  uint   `form:"comment_id"`                      // 父评论 ID
	UserID     uint   `form:"user_id" validate:"required"`     //  发表者id
	Ext        string `form:"ext"`                             // 额外信息存储
	Action     uint   `form:"action"`                          // 行为
}

type CommentParamsTop struct {
	ResourceId uint   `form:"resource_id" validate:"required"` // 评论所关联的资源id
	CommentID  uint   `form:"comment_id"`                      // 父评论 ID
	UserID     uint   `form:"user_id" validate:"required"`     //  发表者id
	Ext        string `form:"ext"`                             // 额外信息存储
}

type CommentParamsHighLight struct {
	ResourceId uint   `form:"resource_id" validate:"required"` // 评论所关联的资源id
	CommentID  uint   `form:"comment_id"`                      // 父评论 ID
	UserID     uint   `form:"user_id" validate:"required"`     //  发表者id
	Ext        string `form:"ext"`                             // 额外信息存储
}
