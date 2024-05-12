package service

import (
	"errors"

	userdao "github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db/model"
	userservice "github.com/longpi1/user-interaction-system/comment/comment-service/model/data"

	"github.com/longpi1/gopkg/libary/utils"
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
