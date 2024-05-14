package data

import (
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"
)

// DeleteComment 删除评论
func DeleteComment(param model.CommentParamsDelete) error {
	// 开启事务
	tx := db.GetClient().Begin()

	// 删除评论内容
	if err := model.DeleteCommentContentWithTx(tx, int(param.CommentID)); err != nil {
		return err
	}

	// 删除评论索引并更新楼层计数
	if err := model.DeleteCommentIndexWithTx(tx, param.CommentID); err != nil {
		return err
	}

	// 递归删除子评论
	if err := model.DeleteChildCommentsWithTx(tx, param.CommentID); err != nil {
		return err
	}

	// 更新用户评论数量
	if err := model.DecreaseCommentCount(tx, param.UserID); err != nil {
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
