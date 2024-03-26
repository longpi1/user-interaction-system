package service

import (
	"comment-service/libary/log"
	"comment-service/model/dao/cache"
	"comment-service/model/dao/db/model"
	"comment-service/model/data"
)

func CommentInteract(param model.CommentParamsInteract) error {
	err := data.CommentInteract(param)
	if err != nil {
		log.Error("评论互动失败：%v", param)
		return err
	}
	// 删除相关评论列表缓存数据
	key := cache.GetCommentListKey(param.ResourceId, param.Pid)
	cache.DeleteCommentListCache(key)
	return nil
}
