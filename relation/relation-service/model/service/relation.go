package service

import (
	"relation-service/libary/log"
	"relation-service/model/dao/db/model"
	"relation-service/model/data"
)

func Relation(params model.RelationParams) error {
	// 校验用户是否被封禁或者黑产等
	err := verify(params)
	if err != nil {
		log.Error("关注校验失败, %v", err)
		return err
	}
	data.Relation(params)

	return nil
}

func ValidateRelationParams(params model.RelationParams) bool {

	return true
}

// verify 校验用户是否被封禁或者黑产等
func verify(params model.RelationParams) error {
	return nil
}
