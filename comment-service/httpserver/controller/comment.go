package controller

import (
	"comment-service/libary/constant"
	"comment-service/libary/log"
	"comment-service/libary/utils"
	"comment-service/model/dao/db/model"
	"comment-service/model/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func AddComment(c *gin.Context) {
	var params model.CommentParamsAdd
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil || !validateParamsAdd(params) {
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
	var params model.CommentParamsList
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	// 校验参数是否正确
	if err != nil || !validateParamsList(params) {
		utils.RespError(c, constant.InvalidParam)
		return
	}

	listResponse, err := service.GetCommentList(params)
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespData(c, "添加评论成功", listResponse)
}

func validateParamsList(params model.CommentParamsList) bool {
	// 参数校验
	return true
}

func DeleteComment(c *gin.Context) {
	// 获取评论ID
	commentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.RespError(c, "无效的评论ID")
		return
	}

	// 验证用户权限
	userID := c.GetInt64("userID") // 假设通过中间件获取了用户ID
	if err := service.VerifyPermission(commentID, userID); err != nil {
		utils.RespError(c, err.Error())
		return
	}

	// 删除评论
	if err := service.DeleteComment(commentID); err != nil {
		log.Error("删除评论失败:", err)
		utils.RespError(c, "删除评论失败")
		return
	}
	utils.RespSuccess(c, "评论删除成功")
}

func CommentDetail(c *gin.Context) {

}
