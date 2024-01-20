package service

import (
	"model-api/model/dao/db/model"
	"model-api/model/data"
)

func SearchLogs(param model.Param) ([]model.Log, error) {
	logs, err := data.SearchLogs(param)
	return logs, err
}

func DeleteLogs(deleteTime int) bool {
	err := data.DeleteLogs(deleteTime)
	if err != nil {
		return false
	}
	return true
}
