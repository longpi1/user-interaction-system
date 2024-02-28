package service

import (
	userdao "comment-service/model/dao/db/model"
	userdata "comment-service/model/data"
	"fmt"
)

func Register(user *userdao.User) error {
	err := userdata.InsertUser(user)
	if err != nil {
		return fmt.Errorf("插入用户信息错误")
	}
	return nil
}
