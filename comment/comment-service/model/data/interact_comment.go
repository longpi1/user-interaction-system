package data

import (
	"errors"

	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"

	"github.com/longpi1/user-interaction-system/comment-service/libary/constant"

	"github.com/longpi1/gopkg/libary/log"
)

func CommentInteract(param model.CommentParamsInteract) error {
	commentIndex, err := model.FindCommentIndexById(int(param.CommentID))
	if err != nil {
		return err
	}
	// 执行互动操作
	switch param.Action {
	case constant.ActionLike:
		// 点赞操作
		commentIndex.LikeCount++
	case constant.ActionDisLike:
		// 点踩操作
		commentIndex.HateCount++
	case constant.ActionReport:
		// 举报操作，需要通知后台审核
		// TODO: 实现举报逻辑
		return nil
	case constant.ActionHighLight:
		commentIndex.IsHighLight = !commentIndex.IsHighLight
	case constant.ActionTop:
		commentIndex.IsPending = !commentIndex.IsPinned
	default:
		log.Info("不存在相关互动动作:", param.Action)
		return errors.New("不存在相关互动动作")
	}

	// 更新评论索引
	if err := model.UpdateCommentIndex(commentIndex); err != nil {
		return err
	}

	return nil
}
