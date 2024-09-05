package model

import (
	"time"

	"github.com/longpi1/user-interaction-system/like-service/model/dao/db"

	"gorm.io/gorm"
)

//	Like 点赞记录表
//
// id BIGINT AUTO_INCREMENT PRIMARY KEY,  -- 自增ID，用于唯一标识每条记录
// uid BIGINT NOT NULL,              -- 用户ID
// business_id BIGINT NOT NULL,           -- 业务ID
// resource_id BIGINT NOT NULL,            -- 被点赞的实体ID
// source VARCHAR(50) NOT NULL,           -- 点赞来源
// op_type   int NOT NULL, -- 操作类型  点赞/点踩
// type   int NOT NULL, -- 点赞类型
// status int NOT NULL,  -- 状态 0未点赞，1已点赞
// is_delete tinyint not null default '0' comment '是否逻辑删除',
// created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 点赞时间
// update_time timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment '更新时间',
// ext     Longtext,  -- 额外信息存储
type Like struct {
	ID         int64     `json:"id" gorm:"primary_key;auto_increment"`
	UID        int64     `json:"uid" gorm:"not null"`
	BusinessID int64     `json:"business_id" gorm:"not null"`
	ResourceID int64     `json:"resource_id" gorm:"not null"`
	Source     string    `json:"source" gorm:"type:varchar(50);not null"`
	OpType     int       `json:"op_type" gorm:"not null"`
	Type       int       `json:"type" gorm:"not null"`
	Status     int       `json:"status" gorm:"not null"`
	IsDelete   bool      `json:"is_delete" gorm:"default:false;comment:'是否逻辑删除'"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime:true;comment:'更新时间'"`
	Ext        string    `json:"ext" gorm:"type:longtext"`
}

// TableName 定义表名
func (Like) TableName() string {
	return "likes"
}

func InsertLike(like *Like) error {
	err := db.GetClient().Create(&like).Error
	return err
}

func InsertLikeWithTx(tx *gorm.DB, like *Like) error {
	err := tx.Create(like).Error
	return err
}

func InsertBatchLike(likes []*Like) error {
	err := db.GetClient().Create(&likes).Error
	return err
}

func DeleteLike(like *Like) error {
	err := db.GetClient().Unscoped().Delete(&like).Error
	return err
}

func DeleteLikeWithTx(tx *gorm.DB) error {
	err := tx.Where().Delete(&Like{}).Error
	return err
}

func FindLikeByLikeId(uint) (Like, error) {
	var like Like
	err := db.GetClient().Where().First(&like).Error
	return like, err
}

func UpdateLike(like *Like) error {
	err := db.GetClient().Updates(&like).Error
	return err
}

func GetLikeList() (likes []Like, err error) {
	tx := db.GetClient()
	//if param.Limit == 0 {
	//	param.Limit = constant.DefaultLimit
	//}
	//err = tx.Order(constant.OrderDescById).Limit(param.Limit).Offset(param.Offset).Find(&likes).Error
	return likes, err
}

//func DeleteLikeByTime(deleteTime int) error {
//	err := db.GetClient().Where(constant.LessThanCreatedAt, deleteTime).Delete(&Like{}).Error
//	return err
//}
