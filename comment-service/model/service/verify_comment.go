package service

import (
	"comment-service/model/dao/db/model"
	"errors"
)

// VerifyPermission 验证用户是否有权限删除评论
func VerifyPermission(commentID, userID uint) error {
	comment, err := model.FindCommentIndexById(int(commentID))
	if err != nil {
		return errors.New("获取对应评论相关信息失败")
	}

	resouceID := comment.ResourceId
	resource, err := model.FindResourceById(int(resouceID))
	if err != nil {
		return errors.New("获取对应资源相关信息失败")
	}
	// 检查用户是否为评论作者或资源发表人
	if comment.UserID != uint(userID) && resource.UserID != uint(userID) {
		return errors.New("没有权限删除该评论")
	}

	return nil
}
