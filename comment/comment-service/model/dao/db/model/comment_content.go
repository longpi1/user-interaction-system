package model

import (
	"github.com/longpi1/user-interaction-system/comment-service/libary/constant"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db"

	"gorm.io/gorm"
)

/*
CommentContent：评论内容表
记录核心评论的内容，避免检索的时候内容过多导致效率低。
*/
type CommentContent struct {
	gorm.Model
	CommentIndex CommentIndex `gorm:"foreignKey:CommentId;comment:'主键'"`
	CommentId    int64        `gorm:"comment:'评论id'"`             // 评论id
	ResourceId   int64        `gorm:"index;comment:'评论所关联的资源id'"` // 评论所关联的资源id
	Content      string       `gorm:"comment:'文本信息'"`             // 文本信息
	ContentMeta  string       `gorm:"comment:'存储一些关键的附属信息'"`      // 存储一些关键的附属信息
	ContentRich  string       `gorm:"comment:'富文本'"`              // 富文本
	Pid          int64        `gorm:"comment:'父评论id'"`            // 父评论 ID
	UserID       int64        `gorm:"index;comment:'发表者id'"`      // 发表者id
	UserName     string       `gorm:"comment:'发表者名称'"`            // 发表者名称
	Ext          string       `gorm:"comment:'额外信息存储'"`           // 额外信息存储
}

// TableName 自定义表名
func (CommentContent) TableName() string {
	return constant.CommentContentTableName
}

func InsertCommentContent(commentContent *CommentContent) error {
	err := db.GetClient().Create(&commentContent).Error
	return err
}

func InsertCommentContentWithTx(tx *gorm.DB, commentContent *CommentContent) error {
	err := tx.Create(commentContent).Error
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

func DeleteCommentContentWithTx(tx *gorm.DB, commentId int) error {
	err := tx.Where(constant.WhereByCommentID, commentId).Delete(&CommentContent{}).Error
	return err
}

func FindCommentContentByCommentId(commentId uint) (CommentContent, error) {
	var commentContent CommentContent
	err := db.GetClient().Where(constant.WhereByCommentID, commentId).First(&commentContent).Error
	return commentContent, err
}

func UpdateCommentContent(commentContent *CommentContent) error {
	err := db.GetClient().Updates(&commentContent).Error
	return err
}

func GetCommentContentList(param CommentParamsList) (CommentContents []CommentContent, err error) {
	tx := db.GetClient()
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
