package model

type RelationParams struct {
	Source     string `json:"source"`      //来源
	UID        int64  `json:"uid"`         // 用户id，也就是发起关注行为的用户id
	Type       string `json:"type"`        // 资源类型
	ResourceID int64  `json:"resource_id"` // 被关注的资源或者人
	Platform   string `json:"platform"`    // 相关的平台
	Status     int    `json:"status"`      // 状态
	OpType     string `json:"op_type"`     // 操作类型
	Ext        string `json:"ext"`         // 额外信息
}

type RelationFansParams struct {
	Source     string `json:"source"`      //来源
	ResourceID int64  `json:"resource_id"` // 被关注的资源或者人
	Type       int    `json:"type"`        // 资源类型
	Platform   int    `json:"platform"`    // 相关的平台
	Status     int    `json:"status"`      // 状态
	Ext        string `json:"ext"`         // 额外信息
}

type RelationFollowingParams struct {
	Source   string `json:"source"`   //来源
	UID      int64  `json:"uid"`      // 发起关注的人
	Type     int    `json:"type"`     // 资源类型
	Platform int    `json:"platform"` // 相关的平台
	Status   int    `json:"status"`   // 状态
	Ext      string `json:"ext"`      // 额外信息
}

type RelationMutualFollowParams struct {
	Source     string `json:"source"`      //来源
	ResourceID int64  `json:"resource_id"` // 被关注的资源或者人
	Type       int    `json:"type"`        // 资源类型
	Platform   int    `json:"platform"`    // 相关的平台
	Status     int    `json:"status"`      // 状态
	Ext        string `json:"ext"`         // 额外信息
}

type RelationCountParams struct {
	ResourceID int64  `json:"resource_id"` // 资源/用户id
	Type       int    `json:"type"`        // 资源类型
	Platform   int    `json:"platform"`    // 相关的平台
	Ext        string `json:"ext"`         // 额外信息
}
