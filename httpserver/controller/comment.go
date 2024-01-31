package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"user-interaction-system/libary/constant"
	"user-interaction-system/libary/utils"
	"user-interaction-system/model/dao/db/model"
	"user-interaction-system/model/service"
)

func AddComment(c *gin.Context) {
	var params model.CommentParamsAdd
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil && !validateParamsAdd(params) {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	err = service.AddComment(params)
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespSuccess(c, "添加评论成功")
}

func validateParamsAdd(params model.CommentParamsAdd) bool {
	// 参数校验
	return true
}

func CommentList(c *gin.Context) {

}

func DeleteComment(c *gin.Context) {

}

func CommentDetail(c *gin.Context) {
	
}
