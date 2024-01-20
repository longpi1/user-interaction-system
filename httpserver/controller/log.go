package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"model-api/libary/constant"
	"model-api/model/dao/db/model"
	"model-api/model/service"
	"net/http"
	"strconv"
)

func GetLogs(c *gin.Context) {
	var logs []model.Log
	var param model.Param
	err := json.NewDecoder(c.Request.Body).Decode(&param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": constant.InvalidParam,
		})
		return
	}
	if logs, err = service.SearchLogs(param); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取日志成功",
		"data":    logs,
	})
}

func DeleteLogs(c *gin.Context) {
	deleteTime, err := strconv.ParseInt(c.Query("delete_threshold_time"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": constant.InvalidParam,
		})
		return
	}

	if !service.DeleteLogs(int(deleteTime)) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "删除日志成功",
	})
}
