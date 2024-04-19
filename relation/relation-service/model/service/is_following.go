package service

import (
	"relation-service/model/dao/db/model"
	"relation-service/model/data"
)

func IsFollowing(params model.RelationIsFollowingParams) (model.RelationIsFollowingResponse, error) {
	return data.IsFollowing(params)
}

func IsFollowingBatch(params model.RelationIsFollowingBatchParams) (model.RelationIsFollowingBatchResponse, error) {
	return data.IsFollowingBatch(params)
}
