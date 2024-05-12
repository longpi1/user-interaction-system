package model

import (
	"github.com/longpi1/user-interaction-system/comment/comment-job/model/dao/db"

	"github.com/longpi1/gopkg/libary/log"
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
