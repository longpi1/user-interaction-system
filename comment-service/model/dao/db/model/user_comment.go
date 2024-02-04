package model

import (
	"gorm.io/gorm"
	"user-interaction-system/libary/constant"
	"user-interaction-system/model/dao/db"
)

/*
UserComment：用户评论相关表
记录评论的索引
同样记录对应的主题，方便后续查询
通过 root 和 parent 记录是否是根评论以及子评论的上级
floor 记录评论层级，也需要更新主题表中的楼层数，
*/
type UserComment struct {
	gorm.Model
	UserID       uint `gorm:"index"` //  发表者id
	PublishCount uint // 发表评论数量
	ReceiveCount uint // 收到评论数量
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
