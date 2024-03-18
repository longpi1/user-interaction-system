package data

import (
	"comment-service/model/dao/db/model"
)

func CommentHighLight(param model.CommentParamsHighLight) error {
	commentIndex, err := model.FindCommentIndexById(int(param.CommentID))
	if err != nil {
		return err
	}
	commentIndex.IsHighLight = param.IsHighLight

	// 更新评论索引
	if err := model.UpdateCommentIndex(&commentIndex); err != nil {
		return err
	}
	return nil
}
