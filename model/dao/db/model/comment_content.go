package model

import (
	"gorm.io/gorm"
	"user-interaction-system/libary/constant"
	"user-interaction-system/model/dao/db"
)

/*
comment_content：评论内容表
记录核心评论的内容，避免检索的时候内容过多导致效率低。
*/
type CommentContent struct {
	gorm.Model
	CommentId     uint   `gorm:"index"` // 评论id
	ResourceId    uint   `gorm:"index"` // 评论所关联的资源id
	ResourceTitle string // 资源的title
	Content       string // 文本信息
	ContentMeta   string // 存储一些关键的附属信息
	ContentRich   string
	Pid           uint   // 父评论 ID
	UserID        uint   `gorm:"index"` //  发表者id
	UserName      string // 发表者名称
	Ext           string // 额外信息存储
}

func InsertCommentContent(commentContent *CommentContent) error {
	err := db.GetClient().Create(&commentContent).Error
	return err
}

func InsertBatchCommentContent(commentContents []*CommentContent) error {
	err := db.GetClient().Create(&commentContents).Error
	return err
}

func DeleteCommentContent(commentContent *CommentContent) error {
	err := db.GetClient().Unscoped().Delete(&commentContent).Error
	return err
}

func FindCommentContentById(id string) (CommentContent, error) {
	var commentContent CommentContent
	err := db.GetClient().Where(constant.WhereByID, id).First(&commentContent).Error
	return commentContent, err
}

func UpdateCommentContent(commentContent *CommentContent) error {
	err := db.GetClient().Updates(&commentContent).Error
	return err
}

func GetCommentContentList(param CommentParamsList) (CommentContents []CommentContent, err error) {
	tx := db.GetClient()
	if param.Type != 0 {
		tx = tx.Where(constant.WhereByType, param.Type)
	}
	if param.Username != "" {
		tx = tx.Where(constant.WhereByUserName, param.Username)
	}
	if param.Content != "" {
		tx = tx.Where(constant.WhereByContent, constant.FuzzySearch+param.Content+constant.FuzzySearch)
	}
	if param.Limit == 0 {
		param.Limit = constant.DefaultLimit
	}
	err = tx.Order(constant.OrderDescById).Limit(param.Limit).Offset(param.Offset).Find(&CommentContents).Error
	return CommentContents, err
}

func DeleteCommentContentByTime(deleteTime int) error {
	err := db.GetClient().Where(constant.LessThanCreatedAt, deleteTime).Delete(&CommentContent{}).Error
	return err
}
