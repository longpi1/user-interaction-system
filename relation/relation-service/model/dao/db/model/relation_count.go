package model

import (
	"relation-service/libary/constant"
	"relation-service/model/dao/db"

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

func InsertRelationCount(relationCount *RelationCount) (uint, error) {
	err := db.GetClient().Create(&relationCount).Error
	return relationCount.ID, err
}

func InsertRelationCountWithTx(tx *gorm.DB, relationCount *RelationCount) (uint, error) {
	err := db.GetClient().Create(&relationCount).Error
	return relationCount.ID, err
}

func InsertBatchRelationCount(RelationCounts []*RelationCount) error {
	err := db.GetClient().Create(&RelationCounts).Error
	return err
}

func DeleteRelationCount(relationCount *RelationCount) error {
	err := db.GetClient().Unscoped().Delete(&relationCount).Error
	return err
}

func DeleteRelationCountWithTx(tx *gorm.DB, id uint) error {
	err := tx.Where(constant.WhereByID, id).Delete(&RelationCount{}).Error
	return err
}

func FindRelationCountById(id int) (RelationCount, error) {
	var relationCount RelationCount
	err := db.GetClient().Where(constant.WhereByID, id).First(&relationCount).Error
	return relationCount, err
}

func UpdateRelationCount(relationCount *RelationCount) error {
	err := db.GetClient().Updates(&relationCount).Error
	return err
}
