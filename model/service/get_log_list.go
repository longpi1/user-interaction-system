package service

import (
	"user-interaction-system/model/dao/db/model"
	"user-interaction-system/model/data"
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
