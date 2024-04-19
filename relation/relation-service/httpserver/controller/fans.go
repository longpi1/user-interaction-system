package controller

import (
	"encoding/json"
	"relation-service/libary/constant"
	"relation-service/libary/utils"
	"relation-service/model/dao/db/model"
	"relation-service/model/service"

	"github.com/gin-gonic/gin"
)

func Fans(c *gin.Context) {
	var params model.RelationFansParams
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil || !validateRelationFansParams(params) {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	fansListResponse, err := service.Fans(params)
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespData(c, "获取粉丝列表成功", fansListResponse)
}

func validateRelationFansParams(params model.RelationFansParams) bool {
	if VerifyTypeAndPlatform(params.Type, params.Platform) {
		return false
	}
	return true
}
