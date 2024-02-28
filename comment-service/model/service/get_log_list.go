package service

import (
	"comment-service/model/dao/db/model"
	"comment-service/model/data"
)

func SearchLogs(param model.LogParam) ([]model.Log, error) {
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
