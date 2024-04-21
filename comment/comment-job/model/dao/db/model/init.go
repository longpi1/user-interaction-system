package model

import (
	"comment-job/libary/log"
	"comment-job/model/dao/db"
)

func InitTable() error {
	// Migrate the schema
	// 注意表的创建顺序，因为有关联字段
	err := db.GetClient().AutoMigrate(&Resource{}, &CommentIndex{}, &UserComment{}, &CommentContent{}, &Permission{})
	if err != nil {
		return err
	}
	log.Info("初始化数据库表成功")
	return nil
}
