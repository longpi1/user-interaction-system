package service

import (
	"relation-service/model/dao/db/model"
	"relation-service/model/data"
)

func Fans(params model.RelationFansParams) (model.RelationFollowingListResponse, error) {
	list, err := data.GetFansList(params)
}
