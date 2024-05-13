package service

import (
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/cache"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"
	"github.com/longpi1/user-interaction-system/comment-service/model/data"

	"github.com/longpi1/gopkg/libary/log"
)

func CommentHighLight(param model.CommentParamsHighLight) error {
	err := data.CommentHighLight(param)
	if err != nil {
		log.Error("删除评论失败：", err.Error())
		return err
	}
	// 删除相关评论列表缓存数据
	key := cache.GetCommentListKey(param.ResourceId, param.Pid)
	cache.DeleteCommentListCache(key)
	return nil
}
