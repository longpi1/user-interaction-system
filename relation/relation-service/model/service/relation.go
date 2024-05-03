package service

import (
	"relation-service/libary/constant"
	"relation-service/model/dao/db/model"
	"relation-service/model/data"

	"github.com/longpi1/gopkg/libary/log"
)

func Relation(params model.RelationParams) (err error) {
	// 校验用户是否被封禁或者黑产等
	err = verify(params)
	if err != nil {
		log.Error("参数校验失败, %v", err)
		return err
	}
	switch params.OpType {
	case constant.TypeFollow:
		err = data.Follow(params)
	case constant.TypeUnFollow:
		err = data.UnFollow(params)
	}
	if err != nil {
		return err
	}

	return nil
}

// verify 校验用户是否被封禁或者黑产等
func verify(params model.RelationParams) error {
	return nil
}
