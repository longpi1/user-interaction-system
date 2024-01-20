package model

import (
	"model-api/libary/constant"
	"model-api/libary/log"
	"model-api/model/dao/db"
)

type User struct {
	Id               int    `json:"id"`
	Username         string `json:"username" gorm:"unique;index" validate:"max=12"`
	Password         string `json:"password" gorm:"not null;" validate:"min=8,max=20"`
	DisplayName      string `json:"display_name" gorm:"index" validate:"max=20"`
	Role             int    `json:"role" gorm:"type:int;default:1"`   // admin, common
	Status           int    `json:"status" gorm:"type:int;default:1"` // enabled, disabled
	Email            string `json:"email" gorm:"index" validate:"max=50"`
	GitHubId         string `json:"github_id" gorm:"column:github_id;index"`
	WeChatId         string `json:"wechat_id" gorm:"column:wechat_id;index"`
	VerificationCode string `json:"verification_code" gorm:"-:all"`                                    // this field is only for Email verification, don't save it to database!
	AccessToken      string `json:"access_token" gorm:"type:char(32);column:access_token;uniqueIndex"` // this token is for system management
	Quota            int    `json:"quota" gorm:"type:int;default:0"`
	UsedQuota        int    `json:"used_quota" gorm:"type:int;default:0;column:used_quota"` // used quota
	RequestCount     int    `json:"request_count" gorm:"type:int;default:0;"`               // request number
	Group            string `json:"group" gorm:"type:varchar(32);default:'default'"`
	AffCode          string `json:"aff_code" gorm:"type:varchar(32);column:aff_code;uniqueIndex"`
	InviterId        int    `json:"inviter_id" gorm:"type:int;column:inviter_id;index"`
}

func GetUserList(limit int, offset int) (user []*User, err error) {
	err = db.GetClient().Limit(limit).Offset(offset).Find(&user).Error
	return user, err
}

func InsertUser(user *User) error {
	err := db.GetClient().Create(&user).Error
	return err
}

func InsertBatchUser(users []*User) error {
	err := db.GetClient().Create(&users).Error
	return err
}

func DeleteUser(user *User) error {
	err := db.GetClient().Unscoped().Delete(&user).Error
	return err
}

func FindUserById(id string) (*User, error) {
	var user *User
	err := db.GetClient().Where(constant.WhereByID, id).First(&user).Error
	return user, err
}

func UpdateUser(user *User) error {
	err := db.GetClient().Updates(&user).Error
	return err
}

func FindUserByName(name string) (*User, error) {
	var user *User
	err := db.GetClient().Where(constant.WhereByName, name).First(&user).Error
	return user, err
}

func IsAdmin(userId int) bool {
	if userId == 0 {
		return false
	}
	var user User
	err := db.GetClient().Where(constant.WhereByID, userId).Select("role").Find(&user).Error
	if err != nil {
		log.Error("no such user " + err.Error())
		return false
	}
	return user.Role >= constant.AdminRole
}
