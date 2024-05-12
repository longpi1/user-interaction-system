package model

import (
	"github.com/longpi1/user-interaction-system/comment/comment-service/libary/constant"
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db"

	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Nid          string `gorm:"index;comment:'资源的nid'"`
	Title        string `gorm:"index;comment:'资源标题'"`
	SubTitle     string `gorm:"index;comment:'资源子标题'"`
	Content      string `gorm:"index;comment:'资源内容'"`
	CommentCount string `gorm:"index;comment:'资源评论数量'"`
	UserID       int64  `gorm:"index;comment:'发表者id'"` // 发表者id
	UserName     string `gorm:"comment:'发表者名称'"`       // 发表者名称
	IP           string `gorm:"comment:'发表者ip'"`       // 发表者ip
	IPArea       string `gorm:"comment:'ip属地'"`        // ip属地
	Status       string `gorm:"index;comment:'资源当前状态'"`
}

type ResourceParamsList struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	Username string `json:"username"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
	OrderBy  string `json:"oderby"`
}

type ResourceParamsAdd struct {
	UserName      string `form:"username"`
	Content       string `form:"content" validate:"required"`
	ResourceId    int64  `form:"content" validate:"required"` // 评论所关联的资源id
	ResourceTitle string `form:"content" validate:"required"` // 资源的title
	ContentMeta   string `form:"content_meta"`                // 存储一些关键的附属信息
	Pid           int64  `form:"pid"`                         // 父评论 ID
	UserID        int64  `form:"user_id" validate:"required"` //  发表者id
	Ext           string `form:"ext"`                         // 额外信息存储
	IP            string `form:"ip"`
	ContentRich   string `form:"content_rich"`
	Type          uint   `form:"type"` // 评论类型，文字评论、图评等"`
}

func TableName() string {
	return constant.ResourceTableName
}

func InsertResource(resource *Resource) error {
	err := db.GetClient().Create(&resource).Error
	return err
}

func InsertBatchResource(resources []*Resource) error {
	err := db.GetClient().Create(&resources).Error
	return err
}

func DeleteResource(resource *Resource) error {
	err := db.GetClient().Unscoped().Delete(&resource).Error
	return err
}

func FindResourceById(id int) (Resource, error) {
	var resource Resource
	err := db.GetClient().Where(constant.WhereByID, id).First(&resource).Error
	return resource, err
}

func UpdateResource(resource *Resource) error {
	err := db.GetClient().Updates(&resource).Error
	return err
}

func GetResourceList(param ResourceParamsList) (resources []Resource, err error) {
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
	err = tx.Order(constant.OrderDescById).Limit(param.Limit).Offset(param.Offset).Find(&resources).Error
	return resources, err
}

func DeleteResourceByTime(deleteTime int) error {
	err := db.GetClient().Where(constant.LessThanCreatedAt, deleteTime).Delete(&Resource{}).Error
	return err
}
