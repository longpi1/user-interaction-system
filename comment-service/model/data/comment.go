package data

import (
	"user-interaction-system/libary/log"
	"user-interaction-system/libary/utils"
	"user-interaction-system/model/dao/db/model"
)

func AddCommentIndex(commentIndex *model.CommentIndex) (uint, error) {
	id, err := model.InsertCommentIndex(commentIndex)
	if err != nil {
		log.Error("评论索引数据插入失败，%v", &commentIndex)
		return 0, err
	}
	return id, err
}

func AddCommentContent(content *model.CommentContent) error {
	err := model.InsertCommentContent(content)
	if err != nil {
		log.Error("评论数据插入失败，%v", &content)
		return err
	}
	return nil
}

// FormatCommentInfo 格式化评论信息
func FormatCommentInfo(param model.CommentParamsAdd) (model.CommentIndex, model.CommentContent) {
	commentIndex := model.CommentIndex{
		UserID:     param.UserID,
		UserName:   param.UserName,
		ResourceId: param.ResourceId,
		IP:         param.IP,
		IPArea:     utils.GetIPArea(param.IP),
		Pid:        param.Pid,
		Type:       param.Type,
	}
	commentContent := model.CommentContent{
		UserID:      param.UserID,
		UserName:    param.UserName,
		ResourceId:  param.ResourceId,
		Pid:         param.Pid,
		Content:     param.Content,
		ContentRich: param.ContentRich,
		ContentMeta: param.ContentMeta,
	}
	return commentIndex, commentContent
}

func GetCommentList(param model.CommentParamsList) ([]model.CommentIndex, error) {
	commentIndexs, err := model.GetCommentIndexList(param)

	return commentIndexs, err
}

func FormatCommentListResponse(comments []model.CommentIndex) model.CommentListResponse {
	return model.CommentListResponse{}
}
