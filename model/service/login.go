package service

import (
	"errors"
	"user-interaction-system/libary/utils"
	userdao "user-interaction-system/model/dao/db/model"
	userservice "user-interaction-system/model/data"
)

func Login(user userdao.User) error {
	password := user.Password
	if user.Username == "" || password == "" {
		return errors.New("用户名或密码为空")
	}
	findUser, err := userservice.FindUserByName(user.Username)
	if err != nil {
		return err
	}
	okay := utils.ValidatePasswordAndHash(password, findUser.Password)
	if !okay {
		return errors.New("用户名或密码错误")
	}
	// token设置todo
	return nil
}
