package data

import "github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db/model"

func CommentTop(param model.CommentParamsTop) error {
	commentIndex, err := model.FindCommentIndexById(int(param.CommentID))
	if err != nil {
		return err
	}
	commentIndex.IsPinned = param.IsPinned

	// 更新评论索引
	if err := model.UpdateCommentIndex(&commentIndex); err != nil {
		return err
	}
	return nil
}
