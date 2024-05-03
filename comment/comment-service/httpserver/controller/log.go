package controller

import (
	"comment-service/libary/constant"
	"comment-service/model/dao/db/model"
	"comment-service/model/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/longpi1/gopkg/libary/utils"
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
