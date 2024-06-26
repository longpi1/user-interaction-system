package service

import (
	"fmt"

	"github.com/longpi1/gopkg/libary/log"
	"github.com/longpi1/gopkg/libary/queue"
	"github.com/longpi1/user-interaction-system/comment-service/libary/conf"
	"github.com/longpi1/user-interaction-system/comment-service/libary/constant"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"
	"github.com/longpi1/user-interaction-system/comment-service/model/data"
)

func AddComment(param model.CommentParamsAdd) error {
	commentIndex, commentContent := data.FormatCommentInfo(param)
	// 审核评论是否通过
	isPass := data.AuditComment(commentContent)
	if !isPass {
		return fmt.Errorf("评论审核未通过")
	}
	// 开始事务
	tx := db.GetClient().Begin()

	// 更新评论索引
	id, err := data.AddCommentIndex(tx, &commentIndex)

	// 写入评论内容
	commentContent.CommentId = int64(id)
	err = data.AddCommentContent(tx, &commentContent)

	// 更新用户评论数量
	var userComment model.UserComment
	userComment.UserID = param.UserID
	err = data.UpdateUserComment(tx, &userComment)

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("添加评论失败")
	}
	// 队列发送
	queueConfig := conf.GetConfig().QueueConfig
	commentInfo := model.FormatCommentInfo(commentIndex, commentContent, constant.AddComment)
	if err := queue.Push(queueConfig.TopicName, commentInfo, queueConfig.Config); err != nil {
		log.Error("队列发送数据失败：", id)
	}
	return nil
}
