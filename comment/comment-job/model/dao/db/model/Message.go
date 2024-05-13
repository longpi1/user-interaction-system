package model

import (
	"github.com/longpi1/user-interaction-system/comment-job/libary/constant"
	"github.com/longpi1/user-interaction-system/comment-job/model/dao/db"

	"github.com/jinzhu/gorm"
)

// Message 消息
type Message struct {
	gorm.Model
	FromId        int64  `gorm:"not null" json:"fromId" form:"fromId"` // 消息的资源id
	ResourceId    uint   `gorm:"comment:'资源id'"`                       // 评论所关联的资源id
	ResourceTitle string `gorm:"comment:'资源标题'"`                       // 资源的title
	FromSource    string `gorm:"comment:'来源'"`
	FromUserID    int64  `gorm:"not null" json:"fromUserId" form:"fromUserId"`                    // 消息发送人id
	ToUserId      int64  `gorm:"not null;index:idx_message_user_id;" json:"userId" form:"userId"` // 用户编号(消息接收人)
	Title         string `gorm:"size:1024" json:"title" form:"title"`                             // 消息标题
	Content       string `gorm:"type:text;not null" json:"content" form:"content"`                // 消息内容
	QuoteContent  string `gorm:"type:text" json:"quoteContent" form:"quoteContent"`               // 引用内容
	Type          int    `gorm:"type:int(11);not null" json:"type" form:"type"`                   // 消息类型
	Platform      string `gorm:"comment:'平台'"`
	ExtraData     string `gorm:"type:text" json:"extraData" form:"extraData"` // 扩展数据
	Owner         string `gorm:"comment:'操作者名称'"`
	Status        int    `gorm:"type:int(11);not null" json:"status" form:"status"` // 状态：0：未读、1：已读
}

// TableName 自定义表名
func (Message) TableName() string {
	return constant.MessageTableName
}

func InsertMessage(message *Message) (uint, error) {
	err := db.GetClient().Create(&message).Error
	return message.ID, err
}

func InsertMessageWithTx(tx *gorm.DB, message *Message) (uint, error) {
	err := tx.Create(&message).Error
	return message.ID, err
}

func InsertBatchMessage(messages []*Message) error {
	err := db.GetClient().Create(&messages).Error
	return err
}

func DeleteMessage(message *Message) error {
	err := db.GetClient().Unscoped().Delete(&message).Error
	return err
}

func DeleteMessageWithTx(tx *gorm.DB, messageID uint) error {
	err := tx.Where(constant.WhereByID, messageID).Delete(&Message{}).Error
	return err
}

func FindMessageById(id int64) (Message, error) {
	var message Message
	err := db.GetClient().Where(constant.WhereByID, id).First(&message).Error
	return message, err
}

func UpdateMessage(message *Message) error {
	err := db.GetClient().Updates(&message).Error
	return err
}

func DeleteMessageByTime(deleteTime int) error {
	err := db.GetClient().Where(constant.LessThanCreatedAt, deleteTime).Delete(&Message{}).Error
	return err
}
