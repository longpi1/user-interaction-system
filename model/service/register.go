package service

import (
	"fmt"
	userdao "user-interaction-system/model/dao/db/model"
	userdata "user-interaction-system/model/data"
)

func Register(user *userdao.User) error {
	err := userdata.InsertUser(user)
	if err != nil {
		return fmt.Errorf("插入用户信息错误")
	}
	return nil
}
