package data

import (
	"errors"
	_ "fmt"
	"user-interaction-system/libary/log"
	"user-interaction-system/libary/utils"
	userdao "user-interaction-system/model/dao/db/model"
)

func InsertUser(user *userdao.User) error {
	var err error
	// 密码加密
	user.Password, err = utils.Password2Hash(user.Password)
	if err != nil {
		log.Error("密码加密失败")
		return errors.New("密码加密失败")
	}
	err = userdao.InsertUser(user)
	if err != nil {
		log.Error("插入用户信息到数据库错误", err.Error())
		return err
	}
	// redis相关操作，todo
	return nil
}

func FindUserByName(userName string) (*userdao.User, error) {
	findUser, err := userdao.FindUserByName(userName)
	if err != nil || findUser == nil {
		return nil, errors.New("用户查找失败")
	}
	// redis相关操作，todo
	return findUser, nil
}
