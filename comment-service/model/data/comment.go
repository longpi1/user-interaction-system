package data

import (
	"comment-service/libary/log"
	"comment-service/libary/utils"
	"comment-service/model/dao/cache"
	"comment-service/model/dao/db"
	"comment-service/model/dao/db/model"

	"gorm.io/gorm"
)

func AddCommentIndex(tx *gorm.DB, commentIndex *model.CommentIndex) (uint, error) {
	id, err := model.InsertCommentIndexWithTx(tx, commentIndex)
	if err != nil {
		log.Error("评论索引数据插入失败，%v", &commentIndex)
		return 0, err
	}
	return id, err
}

func AddCommentContent(tx *gorm.DB, content *model.CommentContent) error {
	err := model.InsertCommentContentWithTx(tx, content)
	if err != nil {
		log.Error("评论数据插入失败，%v", &content)
		return err
	}
	return nil
}

func UpdateUserComment(tx *gorm.DB, userComment *model.UserComment) error {
	if err := tx.Where("user_id = ?", userComment.UserID).First(&userComment).Error; err != nil {
		userComment.PublishCount++
	}
	if err := tx.Save(&userComment).Error; err != nil {
		log.Error("用户评论数据插入失败，%v", &userComment)
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
		PID:        param.Pid,
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

func GetCommentList(param model.CommentParamsList) (model.CommentListResponse, error) {
	// localcache相关操作
	if response, err := cache.GetCommentListFromLocalCache(param); err == nil {
		return response, nil
	}

	// redis相关操作
	if response, err := cache.GetCommentListFromRedisCache(param); err == nil {
		return response, nil
	}

	// 查找评论索引集合
	commentIndexs, err := model.GetCommentIndexList(param)

	var commentResponses []model.CommentResponse
	for _, index := range commentIndexs {
		// 根据index id 查找对应的content
		if commentContent, err := model.FindCommentContentByCommentId(index.ID); err == nil {
			commentResponse := FormatCommentResponse(index, commentContent)
			commentResponses = append(commentResponses, commentResponse)
		}
	}
	listCount, err := model.GetCommentListCount(param)
	commentListResponse := model.CommentListResponse{
		CommentResponses: commentResponses,
		RootReplyCount:   uint(listCount),
	}

	// 将查询结果更新到缓存中
	cache.SetCommentListToLocalCache(param, commentListResponse)
	cache.SetCommentListToRedisCache(param, commentListResponse)

	return commentListResponse, err
}

func FormatCommentResponse(index model.CommentIndex, content model.CommentContent) model.CommentResponse {
	return model.CommentResponse{
		CommentId:   index.ID,
		ResourceId:  index.ResourceId,
		Pid:         index.PID,
		Type:        index.Type,
		Content:     content.Content,
		ContentRich: content.ContentRich,
		ContentMeta: content.ContentMeta,
		Ext:         content.Ext,
		IP:          index.IP,
		IPArea:      index.IPArea,
		State:       index.State,
		LikeCount:   index.LikeCount,
		HateCount:   index.HateCount,
		FloorCount:  index.FloorCount,
		ReplyCount:  index.ReplyCount,
		UserID:      content.UserID,
		UserName:    content.UserName,
		IsCollapsed: index.IsCollapsed,
		IsPending:   index.IsPending,
		IsPinned:    index.IsPinned,
	}
}

// DeleteComment 删除评论
func DeleteComment(commentID int64) error {
	// 开启事务
	tx := db.GetClient().Begin()
	defer tx.Rollback()

	// 删除评论内容
	if err := model.DeleteCommentContentWithTx(tx, uint(commentID)); err != nil {
		return err
	}

	// 删除评论索引并更新楼层计数
	if err := model.DeleteCommentIndexWithTx(tx, uint(commentID)); err != nil {
		return err
	}

	// 递归删除子评论
	if err := model.DeleteChildComments(tx, commentID); err != nil {
		return err
	}

	// 更新用户评论数量
	if err := model.DecreaseCommentCount(tx, commentID); err != nil {
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
