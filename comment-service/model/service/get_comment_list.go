package service

import (
	"fmt"
	"user-interaction-system/model/dao/db/model"
	"user-interaction-system/model/data"
)

func GetCommentList(param model.CommentParamsList) (model.CommentListResponse, error) {
	commentList, err := data.GetCommentList(param)
	if err != nil {
		return model.CommentListResponse{}, fmt.Errorf("获取评论失败")
	}

	return commentList, nil
}
