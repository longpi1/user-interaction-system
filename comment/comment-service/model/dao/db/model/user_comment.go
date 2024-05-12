package model

import (
	"github.com/longpi1/user-interaction-system/comment/comment-service/libary/constant"
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db"

	"gorm.io/gorm"
)

/*
UserComment：用户评论相关表
查看用户发表评论数量，以及收评数量
*/
type UserComment struct {
	gorm.Model
	UserID       int64 `gorm:"index;comment:'发表者id'"` //  发表者id
	PublishCount uint  `gorm:"comment:'发表评论数量'"`      // 发表评论数量
	ReceiveCount uint  `gorm:"comment:'收到评论数量'"`      // 收到评论数量
}

// TableName 自定义表名
func (UserComment) TableName() string {
	return constant.UserCommentTableName
}

func InsertUserComment(userComment *UserComment) error {
	err := db.GetClient().Create(&userComment).Error
	return err
}

func InsertBatchUserComment(userComments []*UserComment) error {
	err := db.GetClient().Create(&userComments).Error
	return err
}

func DeleteUserComment(userComment *UserComment) error {
	err := db.GetClient().Unscoped().Delete(&userComment).Error
	return err
}

func FindUserCommentById(id string) (UserComment, error) {
	var userComment UserComment
	err := db.GetClient().Where(constant.WhereByID, id).First(&userComment).Error
	return userComment, err
}

func UpdateUserComment(userComment *UserComment) error {
	err := db.GetClient().Updates(&userComment).Error
	return err
}

func DeleteUserCommentByTime(deleteTime int) error {
	err := db.GetClient().Where(constant.LessThanCreatedAt, deleteTime).Delete(&UserComment{}).Error
	return err
}

func UpdateUserCommentWithTx(tx *gorm.DB, userComment *UserComment) error {
	err := tx.Create(&userComment).Error
	return err
}

func DecreaseCommentCount(tx *gorm.DB, userID uint) error {
	var userComment UserComment
	if err := tx.First(&userComment).Error; err != nil {
		return err
	}

	// Decrease count by 1
	userComment.PublishCount -= 1

	// Save changes
	if err := tx.Save(&userComment).Error; err != nil {
		return err
	}
	return nil
}
