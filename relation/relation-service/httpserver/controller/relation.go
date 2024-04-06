package controller

import (
	"encoding/json"
	"relation-service/libary/conf"
	"relation-service/libary/constant"
	"relation-service/libary/log"
	"relation-service/libary/utils"
	"relation-service/model/dao/db/model"
	"relation-service/model/service"

	"github.com/gin-gonic/gin"
)

func Relation(c *gin.Context) {
	var params model.RelationParams
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil || !validateRelationParams(params) {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	err = service.Relation(params)
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespSuccess(c, "添加评论成功")
}

func validateRelationParams(params model.RelationParams) bool {
	if params.OpType != constant.TypeFollow && params.OpType != constant.TypeUnFollow {
		log.Error("操作类型错误")
		return false
	}
	mapConfig := conf.GetMapConfig()
	_, hasType := mapConfig.TypeMap[params.Type]
	if !hasType {
		log.Error("类型不存在")
		return false
	}
	_, hasPlatform := mapConfig.PlatformMap[params.Platform]
	if !hasPlatform {
		log.Error("平台不存在")
		return false
	}
	return true
}
