package model

import (
	"github.com/longpi1/user-interaction-system/comment-service/libary/constant"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	UserId           int    `json:"user_id" gorm:"index"`
	Type             int    `json:"type" gorm:"index:idx_created_at_type"`
	Content          string `json:"content"`
	Username         string `json:"username" gorm:"index:index_username_model_name,priority:2;default:''"`
	ModelName        string `json:"model_name" gorm:"index;index:index_username_model_name,priority:1;default:''"`
	PromptTokens     int    `json:"prompt_tokens" gorm:"default:0"`
	CompletionTokens int    `json:"completion_tokens" gorm:"default:0"`
}

type LogParam struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	Username string `json:"username"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
}

// TableName 自定义表名
func (Log) TableName() string {
	return "log"
}

func InsertLog(log *Log) error {
	err := db.GetClient().Create(&log).Error
	return err
}

func InsertBatchLog(logs []*Log) error {
	err := db.GetClient().Create(&logs).Error
	return err
}

func DeleteLog(log *Log) error {
	err := db.GetClient().Unscoped().Delete(&log).Error
	return err
}

func FindLogById(id string) (Log, error) {
	var log Log
	err := db.GetClient().Where(constant.WhereByID, id).First(&log).Error
	return log, err
}

func UpdateLog(log *Log) error {
	err := db.GetClient().Updates(&log).Error
	return err
}

func GetLogList(param LogParam) (logs []Log, err error) {
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
	err = tx.Order(constant.OrderDescById).Limit(param.Limit).Offset(param.Offset).Find(&logs).Error
	return logs, err
}

func DeleteLogByTime(deleteTime int) error {
	err := db.GetClient().Where(constant.LessThanCreatedAt, deleteTime).Delete(&Log{}).Error
	return err
}
