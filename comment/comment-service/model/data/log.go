package data

import (
	"errors"

	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"

	"github.com/longpi1/gopkg/libary/log"
)

func SearchLogs(param model.LogParam) ([]model.Log, error) {
	logs, err := model.GetLogList(param)
	if err != nil {
		log.Error("获取日志失败")
		return nil, errors.New("获取日志失败")
	}
	return logs, err
}

func DeleteLogs(deleteTime int) error {
	err := model.DeleteLogByTime(deleteTime)
	if err != nil {
		log.Error("删除日志失败")
		return errors.New("获取日删除日志失败志失败")
	}
	return err
}
