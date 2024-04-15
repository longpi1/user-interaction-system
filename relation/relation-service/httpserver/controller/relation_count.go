package controller

import (
	"encoding/json"
	"relation-service/libary/constant"
	"relation-service/libary/utils"
	"relation-service/model/dao/db/model"
	"relation-service/model/service"

	"github.com/gin-gonic/gin"
)

func RelationCount(c *gin.Context) {
	var params model.RelationCountParams
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil || !validateRelationCountParams(params) {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	relationCountResponse, err := service.RelationCount(params)
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespData(c, "获取关注数成功", relationCountResponse)
}

func validateRelationCountParams(params model.RelationCountParams) bool {
	if VerifyTypeAndPlatform(params.Type, params.Platform) {
		return false
	}
	return true
}
