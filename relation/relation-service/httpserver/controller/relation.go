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
	if VerifyTypeAndPlatform(utils.ConvertType(params.Type), utils.ConvertPlatform(params.Platform)) {
		return false
	}
	return true
}
