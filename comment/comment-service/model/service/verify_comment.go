package service

import (
	"errors"

	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db/model"
)

// VerifyDeletePermission 验证用户是否有权限删除评论
func VerifyDeletePermission(commentID, userID int64) error {
	comment, err := model.FindCommentIndexById(int(commentID))
	if err != nil {
		return errors.New("获取对应评论相关信息失败")
	}

	resourceID := comment.ResourceID
	resource, err := model.FindResourceById(int(resourceID))
	if err != nil {
		return errors.New("获取对应资源相关信息失败")
	}
	// 检查用户是否为评论作者或资源发表人
	if comment.UserID != userID && resource.UserID != userID {
		return errors.New("权限不足")
	}

	return nil
}

// VerifyHighLightPermission 验证用户是否有权限高亮
func VerifyHighLightPermission(commentID, userID int64) error {
	comment, err := model.FindCommentIndexById(int(commentID))
	if err != nil {
		return errors.New("获取对应评论相关信息失败")
	}

	resourceID := comment.ResourceID
	resource, err := model.FindResourceById(int(resourceID))
	if err != nil {
		return errors.New("获取对应资源相关信息失败")
	}
	// 检查用户是否为资源发表人
	if resource.UserID != userID {
		return errors.New("权限不足")
	}

	return nil
}

// VerifyTopPermission 验证用户是否有权限置顶
func VerifyTopPermission(commentID, userID int64) error {
	comment, err := model.FindCommentIndexById(int(commentID))
	if err != nil {
		return errors.New("获取对应评论相关信息失败")
	}

	resourceID := comment.ResourceID
	resource, err := model.FindResourceById(int(resourceID))
	if err != nil {
		return errors.New("获取对应资源相关信息失败")
	}
	// 检查用户是否为资源发表人
	if resource.UserID != userID {
		return errors.New("权限不足")
	}
	return nil
}
