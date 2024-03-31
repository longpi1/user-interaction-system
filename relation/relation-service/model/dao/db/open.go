package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func OpenMySql(dsn string, gormConfig *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), gormConfig)
}

func OpenPostgreSQL(dsn string, gormConfig *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), gormConfig)
}

func OpenSqlServer(dsn string, gormConfig *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlserver.Open(dsn), gormConfig)
}
