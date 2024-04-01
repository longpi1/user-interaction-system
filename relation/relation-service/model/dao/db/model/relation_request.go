package model

type RelationParams struct {
	Source     int64  `json:"source"`      //来源
	UID        int64  `json:"uid"`         // 用户id，也就是发起关注行为的用户id
	Type       int64  `json:"type"`        // 资源类型
	ResourceID int64  `json:"resource_id"` // 被关注的资源或者人
	Platform   int64  `json:"platform"`    // 相关的平台
	Status     int    `json:"status"`      // 状态
	OpType     string `json:"op_type"`     // 操作类型
	Ext        string `json:"ext"`         // 额外信息
}

type RelationCountParams struct {
	Source     int64  `json:"source"`      //来源
	UID        int64  `json:"uid"`         // 用户id，也就是发起关注行为的用户id
	Type       int64  `json:"type"`        // 资源类型
	ResourceID int64  `json:"resource_id"` // 被关注的资源或者人
	Platform   int64  `json:"platform"`    // 相关的平台
	Status     int    `json:"status"`      // 状态
	OpType     string `json:"op_type"`     // 操作类型
	Ext        string `json:"ext"`         // 额外信息
}
