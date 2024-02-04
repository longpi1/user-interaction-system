package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"strconv"
	"user-interaction-system/libary/constant"
	"user-interaction-system/libary/utils"
	"user-interaction-system/model/dao/db/model"
	"user-interaction-system/model/service"
)

func GetLogs(c *gin.Context) {
	var logs []model.Log
	var param model.LogParam
	err := json.NewDecoder(c.Request.Body).Decode(&param)
	if err != nil {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	if logs, err = service.SearchLogs(param); err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespData(c, "获取日志成功", logs)
}

func DeleteLogs(c *gin.Context) {
	deleteTime, err := strconv.ParseInt(c.Query("delete_threshold_time"), 10, 64)
	if err != nil {
		utils.RespError(c, constant.InvalidParam)
		return
	}

	if !service.DeleteLogs(int(deleteTime)) {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespSuccess(c, "删除日志成功")
}
