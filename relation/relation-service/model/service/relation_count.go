package service

import (
	"relation-service/model/dao/db/model"
	"relation-service/model/data"
)

func RelationCount(params model.RelationCountParams) (model.RelationCountResponse, error) {
	relationCount, err := data.GetRelationCount(params)
	return relationCount, err
}
