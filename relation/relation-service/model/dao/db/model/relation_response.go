package model

type RelationResponse struct {
	ID           int64  `json:"id"`
	Source       string `json:"source"` //来源
	UID          int64  `json:"uid"`    // 用户id，也就是发起关注行为的用户id
	UserName     string `json:"userName"`
	Type         string `json:"type"`        // 资源类型
	ResourceID   int64  `json:"resource_id"` // 被关注的资源或者人
	ResourceName string `json:"resource_name"`
	Platform     string `json:"platform"` // 相关的平台
	Status       int    `json:"status"`   // 状态
	Ext          string `json:"ext"`      // 额外信息
}

type RelationListResponse struct {
	RelationResponse []RelationResponse
	SubscribeNum     int // 关注数
	FansNum          int // 粉丝数
}
