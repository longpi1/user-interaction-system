package controller

import (
	"encoding/json"
	"relation-service/libary/constant"
	"relation-service/model/dao/db/model"
	"relation-service/model/service"

	"github.com/longpi1/gopkg/libary/utils"

	"github.com/longpi1/gopkg/libary/log"

	"github.com/gin-gonic/gin"
)

func IsFollowing(c *gin.Context) {
	var params model.RelationIsFollowingParams
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil || !validateRelationIsFollowingParams(params) {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	response, err := service.IsFollowing(params)
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespData(c, "查询关注关系成功", response)
}

func validateRelationIsFollowingParams(params model.RelationIsFollowingParams) bool {
	if VerifyTypeAndPlatform(utils.ConvertType(params.Type), utils.ConvertPlatform(params.Platform)) {
		return false
	}
	return true
}

func IsFollowingBatch(c *gin.Context) {
	var params model.RelationIsFollowingBatchParams
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil || !validateRelationIsFollowingBatchParams(params) {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	response, err := service.IsFollowingBatch(params)
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespData(c, "查询关注关系成功", response)
}

func validateRelationIsFollowingBatchParams(params model.RelationIsFollowingBatchParams) bool {
	if len(params.ResourceIDs) == 0 {
		log.Error("参数被关注id为空")
		return false
	}
	if VerifyTypeAndPlatform(utils.ConvertType(params.Type), utils.ConvertPlatform(params.Platform)) {
		return false
	}
	return true
}
