package model

import (
	"user-interaction-system/libary/log"
	"user-interaction-system/model/dao/db"
)

func InitTable() error {
	// Migrate the schema
	// 注意表的创建顺序，因为有关联字段
	err := db.GetClient().AutoMigrate(&CommentIndex{}, &UserComment{}, &User{}, &CommentContent{}, &Permission{}, &Log{})
	if err != nil {
		return err
	}
	log.Info("初始化数据库表成功")
	return nil
}
