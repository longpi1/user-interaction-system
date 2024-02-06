package service

import (
	"fmt"
	"user-interaction-system/model/dao/db/model"
	"user-interaction-system/model/data"
)

func GetCommentList(param model.CommentParamsList) (model.CommentListResponse, error) {
	commentList, err := data.GetCommentList(param)
	if err != nil {
		return model.CommentListResponse{}, fmt.Errorf("添加评论失败")
	}
	listResponse := data.FormatCommentListResponse(commentList)

	return listResponse, nil
}
