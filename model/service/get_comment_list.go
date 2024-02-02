package service

import (
	"fmt"
	"user-interaction-system/model/dao/db/model"
)

func GetCommentList(param model.CommentParamsList) (model.CommentListResponse, error) {

	if !isPass {
		return nil, fmt.Errorf("评论审核未通过")
	}

	if err != nil {
		return nil, fmt.Errorf("添加评论失败")
	}

	return nil, nil
}
