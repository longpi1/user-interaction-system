package model

import (
	"relation-service/libary/constant"

	"gorm.io/gorm"
)

// RelationCount 用户关注数/粉丝数表
type RelationCount struct {
	gorm.Model
	ResourceId  int64  `json:"resource_id"`  // 资源/用户id
	FansCount   int64  `json:"fans_count"`   // 粉丝数
	FollowCount int64  `json:"follow_count"` // 关注数
	Platform    int64  `json:"platform"`     // 相关的平台
	Type        int64  `json:"type"`         // 资源类型
	Ext         string `json:"ext"`          // 额外信息`
}

// TableName 自定义表名
func (RelationCount) TableName() string {
	return constant.RelationCountTableName
}
