package service

import (
	"comment-service/libary/log"
	"comment-service/model/dao/cache"
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
	// 删除相关评论列表缓存数据
	key := cache.GetCommentListKey(param.ResourceId, param.Pid)
	cache.DeleteCommentListCache(key)
	return nil
}
