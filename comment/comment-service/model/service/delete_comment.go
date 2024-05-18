package service

import (
	"github.com/longpi1/gopkg/libary/queue"
	"github.com/longpi1/user-interaction-system/comment-service/libary/conf"
	"github.com/longpi1/user-interaction-system/comment-service/libary/constant"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"
	"github.com/longpi1/user-interaction-system/comment-service/model/data"

	"github.com/longpi1/user-interaction-system/comment-service/model/dao/cache"

	"github.com/longpi1/gopkg/libary/log"
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

	// 队列发送
	queueConfig := conf.GetConfig().QueueConfig
	commentInfo := model.CommentInfo{
		CommentId:  param.CommentID,
		ResourceId: param.ResourceId,
		UserID:     param.UserID,
		Ext:        param.Ext,
		Pid:        param.Pid,
		Type:       constant.DeleteComment,
	}
	if err := queue.Push(queueConfig.TopicName, commentInfo, queueConfig.Config); err != nil {
		log.Error("队列发送数据失败：", param.CommentID)
	}
	return nil
}
