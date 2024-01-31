package service

import (
	"fmt"
	"user-interaction-system/model/dao/db/model"
	"user-interaction-system/model/data"
)

func AddComment(param model.CommentParamsAdd) error {
	commentIndex, commentContent := data.FormatCommentInfo(param)
	// 审核评论是否通过
	isPass := data.AuditComment(commentContent)
	if !isPass {
		return fmt.Errorf("评论审核未通过")
	}

	id, err := data.AddCommentIndex(&commentIndex)
	commentContent.CommentId = id
	err = data.AddCommentContent(&commentContent)
	if err != nil {
		return fmt.Errorf("添加评论失败")
	}

	return nil
}
