package data

import "comment-service/model/dao/db/model"

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
