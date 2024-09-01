package model

// Count 点赞数表
// counts
//
//	business_id BIGINT NOT NULL,           -- 业务ID
//	resource_id BIGINT NOT NULL,            -- 实体ID
//	likes_count INT DEFAULT 0,             -- 点赞数
//	dislikes_count INT DEFAULT 0,          -- 点踩数
type Count struct {
	BusinessID    int64 `json:"business_id" gorm:"primary_key;not null"`
	ResourceID    int64 `json:"resource_id" gorm:"primary_key;not null"`
	LikesCount    int   `json:"likes_count" gorm:"default:0"`
	DislikesCount int   `json:"dislikes_count" gorm:"default:0"`
}

// TableName 定义表名
func (Count) TableName() string {
	return "counts"
}
