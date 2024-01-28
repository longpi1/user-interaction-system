package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"user-interaction-system/libary/constant"
	"user-interaction-system/model/dao/db/model"
	"user-interaction-system/model/service"
)

func AddComment(c *gin.Context) {
	var params model.CommentParamsAdd
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil && !validateParams(params) {
		c.JSON(http.StatusOK, gin.H{
			"message": constant.InvalidParam,
			"success": false,
		})
		return
	}
	err = service.AddComment(params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": constant.InvalidParam,
			"success": false,
		})
		return
	}

}

func validateParams(params model.CommentParamsAdd) bool {
	// 参数校验
	return true
}
