package controller

import (
	"encoding/json"
	"relation-service/libary/constant"
	"relation-service/model/dao/db/model"
	"relation-service/model/service"

	"github.com/longpi1/gopkg/libary/utils"

	"github.com/gin-gonic/gin"
)

func Following(c *gin.Context) {
	var params model.RelationFollowingParams
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil || !validateRelationFollowingParams(params) {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	followingList, err := service.Following(params)
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespData(c, "获取关注列表成功", followingList)
}

func validateRelationFollowingParams(params model.RelationFollowingParams) bool {
	if VerifyTypeAndPlatform(params.Type, params.Platform) {
		return false
	}
	return true
}
