package db

import (
	"gorm.io/gorm"
	"model-api/libary/conf"
	"model-api/libary/log"
	"model-api/model/dao/db/model"
)

var db DB

type DB struct {
	client *gorm.DB
}

func GetClient() *gorm.DB{
	return db.client
}

func NewClient(config conf.DBConf) (*gorm.DB, error){
	gormConfig := &gorm.Config{
		DryRun: true,
	}
	var err error
	switch config.Type {
	case conf.TypeMySql:
		db.client, err = OpenMySql(config.Dsn, gormConfig)
	case conf.TypePostgreSQL:
		db.client, err = OpenPostgreSQL(config.Dsn, gormConfig)
	case conf.TypeMSSQL:
		db.client, err = OpenSqlServer(config.Dsn, gormConfig)
	}
	return db.client, err
}

func InitTable() error{
	// Migrate the schema
	// 注意表的创建顺序，因为有关联字段
	err := db.client.AutoMigrate(model.Model{}, model.Permission{})
	if err != nil {
		return err
	}
	log.Info("初始化数据库表成功")
	return nil
}

