package service

import (
	"fmt"

	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db/model"
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/data"
)

func GetCommentList(param model.CommentParamsList) (model.CommentListResponse, error) {
	commentList, err := data.GetCommentList(param)
	if err != nil {
		return model.CommentListResponse{}, fmt.Errorf("获取评论失败")
	}

	return commentList, nil
}
