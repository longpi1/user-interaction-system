package service

import (
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/cache"
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db/model"
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/data"

	"github.com/longpi1/gopkg/libary/log"
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
