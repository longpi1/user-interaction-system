package model

type RelationResponse struct {
	ID           int64  `json:"id"`
	Source       string `json:"source"` //来源
	UID          int64  `json:"uid"`    // 用户id，也就是发起关注行为的用户id
	UserName     string `json:"userName"`
	Type         int    `json:"type"`        // 资源类型
	ResourceID   int64  `json:"resource_id"` // 被关注的资源或者人
	ResourceName string `json:"resource_name"`
	Platform     int    `json:"platform"` // 相关的平台
	Status       int    `json:"status"`   // 状态
	Ext          string `json:"ext"`      // 额外信息
}

type RelationFansListResponse struct {
	RelationResponse []RelationResponse
	FansCount        int // 粉丝数
}

type RelationFollowingListResponse struct {
	RelationResponse []RelationResponse
	FollowingCount   int // 关注数
}

type RelationCountResponse struct {
	ResourceID  int64 `json:"uid"` // 用户/资源id
	FollowCount int   // 关注数
	FansCount   int   // 粉丝数
	Platform    int   `json:"platform"` // 相关的平台
	Type        int   `json:"type"`     // 资源类型
}
