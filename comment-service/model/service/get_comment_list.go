package service

import (
	"comment-service/model/dao/db/model"
	"comment-service/model/data"
	"fmt"
)

func GetCommentList(param model.CommentParamsList) (model.CommentListResponse, error) {
	commentList, err := data.GetCommentList(param)
	if err != nil {
		return model.CommentListResponse{}, fmt.Errorf("获取评论失败")
	}

	return commentList, nil
}
