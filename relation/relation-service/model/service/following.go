package service

import (
	"relation-service/model/dao/db/model"
	"relation-service/model/data"
)

func Following(params model.RelationFollowingParams) (model.RelationFollowingListResponse, error) {
	followingList, err := data.GetFollowingList(params)
	return followingList, err
}
