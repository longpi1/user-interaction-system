package model

import (
	"relation-service/libary/constant"
	"relation-service/model/dao/db"

	"gorm.io/gorm"
)

// RelationCount 用户关注数/粉丝数表
type RelationCount struct {
	gorm.Model
	ResourceId  int64  `gorm:"uniqueIndex:idx_relation_count_resource_id_platform_type;comment:'资源/用户id'"` // 资源/用户id
	FansCount   int64  `gorm:"comment:'粉丝数'"`                                                              // 粉丝数
	FollowCount int64  `gorm:"comment:'关注数'"`                                                              // 关注数
	Platform    int64  `gorm:"uniqueIndex:idx_relation_count_resource_id_platform_type;comment:'相关的平台d'"`  // 相关的平台
	Type        int64  `gorm:"uniqueIndex:idx_relation_count_resource_id_platform_type;comment:'资源类型'"`    // 资源类型
	Ext         string `gorm:"comment:'额外信息'"`                                                             // 额外信息`
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

func FindRelationCountByParams(params RelationCountParams) (RelationCount, error) {
	var relationCount RelationCount
	client := db.GetClient()
	if params.Platform >= 0 {
		client.Where(constant.WhereByPlatform, params.Platform)
	}
	if params.Type >= 0 {
		client.Where(constant.WhereByType, params.Type)
	}
	err := db.GetClient().Where(constant.WhereByResourceID, params.ResourceID).First(&relationCount).Error
	return relationCount, err
}

func UpdateRelationCount(relationCount *RelationCount) error {
	err := db.GetClient().Updates(&relationCount).Error
	return err
}
