package model

import (
	"gorm.io/gorm"
	"model-api/model/dao/db"
)

const(
	WhereByID = "id = ?"
)

type Model struct {
	gorm.Model
	Id         string                  `gorm:"index;size:255"`
	Description string 					`gorm:"index;size:255"`
	OwnedBy    string                  `json:"owned_by"`
}

func GetModelList(limit int, offset int) (models []*Model, err error)  {
	err = db.GetClient().Order("id desc").Limit(limit).Offset(offset).Find(&models).Error
	return models, err
}


func InsertModel(model *Model) error {
	err := db.GetClient().Create(model).Error
	return err
}

func InsertBatchModel(models []*Model) error {
	err := db.GetClient().Create(models).Error
	return err
}

func DeleteModelById(model *Model) error {
	err := db.GetClient().Unscoped().Delete(model).Error
	return err
}

func FindModelById(id string) (Model, error) {
	var model Model
	err := db.GetClient().Where(WhereByID, id).First(&model).Error
	return model, err
}