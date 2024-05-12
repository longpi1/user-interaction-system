package data

import (
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db/model"
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
