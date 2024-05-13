package service

import (
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"
	"github.com/longpi1/user-interaction-system/comment-service/model/data"
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
