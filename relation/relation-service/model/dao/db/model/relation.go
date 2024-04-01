package model

import (
	"relation-service/libary/constant"
	"relation-service/model/dao/db"

	"gorm.io/gorm"
)

// Relation 关注信息表
type Relation struct {
	gorm.Model
	Source     int64   `gorm:"uniqueIndex:idx_relation_platform_source_uid;comment:'来源'"`     //来源
	UID        int64  `gorm:"uniqueIndex:idx_relation_platform_source_uid;comment:'用户id，也就是发起关注行为的用户id'"`         // 用户id，也就是发起关注行为的用户id
	ResourceID int64    `gorm:"index;comment:'被关注的资源或者人id'"` // 被关注的资源或者人id
	Platform   int64    `gorm:"uniqueIndex:idx_relation_platform_source_uid;comment:'相关的平台'"`     // 相关的平台
	Status     int     `gorm:"comment:'状态'"`    // 状态
	Type     int     `gorm:"comment:'类型'"`    // 类型
	Ext        string   `gorm:"comment:'额外信息'"`          // 额外信息
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
	if err := tx.Where("resource_id = ? AND pid = ?", relation.ResourceId, relation.PID).Last(&relation).Error; err != nil {
		relation.FloorCount = 1
	} else {
		relation.FloorCount++
	}
	err := db.GetClient().Create(&relation).Error
	return relation.ID, err
}

func InsertBatchRelation(Relations []*Relation) error {
	err := db.GetClient().Create(&Relations).Error
	return err
}

func DeleteRelation(relation *Relation) error {
	err := db.GetClient().Unscoped().Delete(&relation).Error
	return err
}

func DeleteRelationWithTx(tx *gorm.DB, commentID uint) error {
	err := tx.Where(constant.WhereByID, commentID).Delete(&Relation{}).Error
	return err
}

func FindRelationById(id int) (Relation, error) {
	var relation Relation
	err := db.GetClient().Where(constant., id).First(&relation).Error
	return relation, err
}

func UpdateRelation(relation *Relation) error {
	err := db.GetClient().Updates(&relation).Error
	return err
}

func GetRelationList(param CommentParamsList) (Relations []Relation, err error) {
	tx := db.GetClient()
	if param.ResourceId != 0 {
		tx = tx.Where(constant.WhereByResourceID, param.ResourceId)
	}
	if param.Pid != 0 {
		tx = tx.Where(constant.WhereByPID, param.Pid)
	}
	if param.Limit == 0 {
		param.Limit = constant.DefaultLimit
	}
	err = tx.Order(constant.OrderDescById).Limit(param.Limit).Offset(param.Offset).Find(&Relations).Error
	return Relations, err
}

func GetCommentListCount(param CommentParamsList) (count int64, err error) {
	tx := db.GetClient()
	if param.ResourceId != 0 {
		tx = tx.Where(constant.WhereByResourceID, param.ResourceId)
	}
	if param.Pid != 0 {
		tx = tx.Where(constant.WhereByPID, param.Pid)
	}
	err = tx.Count(&count).Error
	return count, err
}

func DeleteRelationByTime(deleteTime int) error {
	err := db.GetClient().Where(constant.LessThanCreatedAt, deleteTime).Delete(&Relation{}).Error
	return err
}

// DeleteChildCommentsWithTx 递归删除子评论
func DeleteChildCommentsWithTx(tx *gorm.DB, commentID uint) error {
	// 查询当前评论的所有子评论
	var childComments []Relation
	if err := tx.Where("pid = ?", commentID).Find(&childComments).Error; err != nil {
		return err
	}

	// 遍历子评论
	for _, childComment := range childComments {
		// 删除子评论的内容
		if err := DeleteCommentContentWithTx(tx, childComment.ID); err != nil {
			return err
		}

		// 删除子评论的索引
		if err := DeleteRelationWithTx(tx, childComment.ID); err != nil {
			return err
		}

		// 更新用户评论数量
		if err := DecreaseCommentCount(tx, childComment.UserID); err != nil {
			return err
		}

		// 递归删除子评论的子评论
		if err := DeleteChildCommentsWithTx(tx, childComment.ID); err != nil {
			return err
		}
	}

	return nil
}

