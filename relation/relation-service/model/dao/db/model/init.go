package model

import (
	"relation-service/libary/log"
	"relation-service/model/dao/db"
)

func InitTable() error {
	// Migrate the schema
	// 注意表的创建顺序，因为有关联字段
	err := db.GetClient().AutoMigrate(&Relation{}, &RelationCount{})
	if err != nil {
		return err
	}
	log.Info("初始化数据库表成功")
	return nil
}
