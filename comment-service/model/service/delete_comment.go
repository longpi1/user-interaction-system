package service

import (
	"comment-service/libary/log"
	"comment-service/model/dao/cache"
	"comment-service/model/data"
)

// DeleteComment 删除评论
func DeleteComment(commentID int64) error {
	err := data.DeleteComment(commentID)
	if err != nil {
		log.Error("删除评论失败：", err.Error())
		return err
	}
	// 删除相关评论列表缓存数据
	cache.DeleteCommentListCache(param)
	return nil
}
