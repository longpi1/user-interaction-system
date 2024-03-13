package controller

import (
	"comment-service/libary/constant"
	"comment-service/libary/log"
	"comment-service/libary/utils"
	"comment-service/model/dao/db/model"
	"comment-service/model/service"

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
	var param model.CommentParamsDelete
	err := json.NewDecoder(c.Request.Body).Decode(&param)
	// 校验参数是否正确
	if err != nil || !validateParamsDelete(param) {
		utils.RespError(c, constant.InvalidParam)
		return
	}

	// 验证用户权限
	if err := service.VerifyPermission(param.CommentID, param.UserID); err != nil {
		utils.RespError(c, err.Error())
		return
	}

	// 删除评论
	if err := service.DeleteComment(param); err != nil {
		log.Error("删除评论失败:", err)
		utils.RespError(c, "删除评论失败")
		return
	}
	utils.RespSuccess(c, "评论删除成功")
}

func validateParamsDelete(params model.CommentParamsDelete) bool {
	// 参数校验
	return true
}

func CommentDetail(c *gin.Context) {

}

// CommentInteract 用户相关互动行为，点赞、点踩、举报
func CommentInteract(c *gin.Context) {
	var param model.CommentParamsInteract
	err := json.NewDecoder(c.Request.Body).Decode(&param)
	// 校验参数是否正确
	if err != nil || !validateParamsInteract(param) {
		utils.RespError(c, constant.InvalidParam)
		return
	}

	if err := service.; err != nil {
		log.Error("评论失败:", err)
		utils.RespError(c, "评论失败")
		return
	}
	utils.RespSuccess(c, "评论成功")
}

func validateParamsInteract(params model.CommentParamsInteract) bool {
	// 参数校验
	return true
}

// CommentTop 置顶某一个评论
func CommentTop(c *gin.Context) {

}

func validateParamsTop(params model.CommentParamsTop) bool {
	// 参数校验
	return true
}

// CommentHighLight 高亮某一个评论
func CommentHighLight(c *gin.Context) {

}


func validateParamsHighLight(params model.CommentParamsHighLight) bool {
	// 参数校验
	return true
}