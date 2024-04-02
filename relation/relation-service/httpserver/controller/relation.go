package controller

import (
	"encoding/json"
	"relation-service/libary/constant"
	"relation-service/libary/utils"
	"relation-service/model/dao/db/model"
	"relation-service/model/service"

	"github.com/gin-gonic/gin"
)

func Relation(c *gin.Context) {
	var params model.RelationParams
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil || !service.ValidateRelationParams(params) {
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
