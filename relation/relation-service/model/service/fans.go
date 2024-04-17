package service

import (
	"relation-service/model/dao/db/model"
	"relation-service/model/data"
)

func Fans(params model.RelationFansParams) (model.RelationFansListResponse, error) {
	list, err := data.GetFansList(params)
	return list, err
}
