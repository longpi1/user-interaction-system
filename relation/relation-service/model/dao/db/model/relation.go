package model

import (
	"relation-service/libary/constant"
	"relation-service/model/dao/db"

	"gorm.io/gorm"
)

// Relation 关注信息表
type Relation struct {
	gorm.Model
	Source     string `gorm:"size:255;uniqueIndex:idx_relation_platform_source_uid;comment:'来源'"`         //来源
	UID        int64  `gorm:"uniqueIndex:idx_relation_platform_source_uid;comment:'用户id，也就是发起关注行为的用户id'"` // 用户id，也就是发起关注行为的用户id
	ResourceID int64  `gorm:"index;comment:'被关注的资源或者人id'"`                                                // 被关注的资源或者人id
	Platform   int    `gorm:"uniqueIndex:idx_relation_platform_source_uid;comment:'相关的平台'"`               // 相关的平台
	Status     int    `gorm:"comment:'状态'"`                                                               // 状态
	Type       int    `gorm:"comment:'类型'"`                                                               // 类型
	Ext        string `gorm:"comment:'额外信息'"`                                                             // 额外信息
}

// TableName 自定义表名
func (Relation) TableName() string {
	return constant.RelationTableName
}

func InsertRelation(relation *Relation) (uint, error) {
	err := db.GetClient().Create(&relation).Error
	return relation.ID, err
}

func InsertRelationWithTx(tx *gorm.DB, relation *Relation) (uint, error) {
	err := db.GetClient().Create(&relation).Error
	return relation.ID, err
}

func InsertBatchRelation(relations []*Relation) error {
	err := db.GetClient().Create(&relations).Error
	return err
}

func DeleteRelation(relation *Relation) error {
	err := db.GetClient().Unscoped().Delete(&relation).Error
	return err
}

func DeleteRelationWithTx(tx *gorm.DB, uid int64, resourceID int64) error {
	err := tx.Where(constant.WhereByUserID, uid).Where(constant.WhereByResourceID, resourceID).Unscoped().Delete(&Relation{}).Error
	return err
}

func FindRelationById(id int) (Relation, error) {
	var relation Relation
	err := db.GetClient().Where(constant.WhereByID, id).First(&relation).Error
	return relation, err
}

func UpdateRelation(relation *Relation) error {
	err := db.GetClient().Updates(&relation).Error
	return err
}
