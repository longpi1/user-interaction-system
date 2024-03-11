package service

import (
	"comment-service/libary/log"
	"comment-service/model/dao/db/model"
	"comment-service/model/data"
)

// DeleteComment 删除评论
func DeleteComment(param model.CommentParamsDelete) error {
	err := data.DeleteComment(param)
	if err != nil {
		log.Error("删除评论失败：", err.Error())
		return err
	}

	return nil
}
