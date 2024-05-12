package model

type CommentParamsList struct {
	ResourceId int64  `form:"content" validate:"required"` // 评论所关联的资源id
	Pid        int64  `form:"pid"`                         // 父评论 ID
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	OrderBy    string `json:"oderby"`
}

type CommentParamsAdd struct {
	UserName      string `form:"username"`
	Content       string `form:"content" validate:"required"`
	ResourceId    int64  `form:"resource_id" validate:"required"`    // 评论所关联的资源id
	ResourceTitle string `form:"resource_title" validate:"required"` // 资源的title
	ContentMeta   string `form:"content_meta"`                       // 存储一些关键的附属信息
	Pid           int64  `form:"pid"`                                // 父评论 ID
	UserID        int64  `form:"user_id" validate:"required"`        //  发表者id
	Ext           string `form:"ext"`                                // 额外信息存储
	IP            string `form:"ip"`
	ContentRich   string `form:"content_rich"`
	Type          uint   `form:"type"` // 评论类型，文字评论、图评等"`
}

type CommentParamsDelete struct {
	ResourceId int64  `form:"resource_id" validate:"required"` // 评论所关联的资源id
	CommentID  int64  `form:"comment_id" validate:"required"`  // 父评论 ID
	UserID     int64  `form:"user_id" validate:"required"`     //  发表者id
	Ext        string `form:"ext"`                             // 额外信息存储
	Pid        int64  `form:"pid"`                             // 父评论 ID
}

type CommentParamsInteract struct {
	ResourceId int64  `form:"resource_id" validate:"required"`               // 评论所关联的资源id
	CommentID  int64  `form:"comment_id" validate:"required"`                // 父评论 ID
	UserID     int64  `form:"user_id" validate:"required"`                   //  发表者id
	Ext        string `form:"ext"`                                           // 额外信息存储
	Action     int64  `form:"action" validate:"required"validate:"required"` // 行为
	Pid        int64  `form:"pid" validate:"required"`                       // 父评论 ID
}

type CommentParamsTop struct {
	ResourceId int64  `form:"resource_id" validate:"required"` // 评论所关联的资源id
	CommentID  int64  `form:"comment_id" validate:"required"`  // 父评论 ID
	UserID     int64  `form:"user_id" validate:"required"`     //  发表者id
	Ext        string `form:"ext"`                             // 额外信息存储
	IsPinned   bool   `form:"is_pinned" validate:"required"`   // 是否置顶
	Pid        int64  `form:"pid" validate:"required"`         // 父评论 ID
}

type CommentParamsHighLight struct {
	ResourceId  int64  `form:"resource_id" validate:"required"`  // 评论所关联的资源id
	CommentID   int64  `form:"comment_id" validate:"required"`   // 父评论 ID
	UserID      int64  `form:"user_id" validate:"required"`      //  发表者id
	Ext         string `form:"ext"`                              // 额外信息存储
	IsHighLight bool   `form:"is_highlight" validate:"required"` // 是否高亮
	Pid         int64  `form:"pid" validate:"required"`          // 父评论 ID
}
