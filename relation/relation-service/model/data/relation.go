package data

import "relation-service/model/dao/db/model"

func Relation(params model.RelationParams) error {

	return nil
}

func formatRelation(params model.RelationParams) model.Relation {
	relation := model.Relation{
		Type: params.Type,
	}
	return relation
}
