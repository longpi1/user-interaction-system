package service

import (
	"fmt"

	userdao "github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"
	userdata "github.com/longpi1/user-interaction-system/comment-service/model/data"
)

func Register(user *userdao.User) error {
	err := userdata.InsertUser(user)
	if err != nil {
		return fmt.Errorf("插入用户信息错误")
	}
	return nil
}
