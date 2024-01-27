package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"user-interaction-system/libary/constant"
	"user-interaction-system/model/dao/db/model"
)

func AddComment(c *gin.Context) {
	var params model.CommenParamsAdd
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": constant.InvalidParam,
			"success": false,
		})
		return
	}
	ValidateParams(params)
}

func ValidateParams(params model.CommenParamsAdd) {

}
