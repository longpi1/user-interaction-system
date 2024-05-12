package data

import (
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/cache"
	"github.com/longpi1/user-interaction-system/comment/comment-service/model/dao/db/model"
)

func GetCommentList(param model.CommentParamsList) (model.CommentListResponse, error) {
	// 将查询结果更新到缓存中
	key := cache.GetCommentListKey(param.ResourceId, param.Pid)
	// localcache相关操作
	if response, err := cache.GetCommentListFromLocalCache(key); err == nil {
		return response, nil
	}

	// redis相关操作
	if response, err := cache.GetCommentListFromRedisCache(key); err == nil {
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

	cache.SetCommentListToLocalCache(key, commentListResponse)
	cache.SetCommentListToRedisCache(key, commentListResponse)

	return commentListResponse, err
}

func FormatCommentResponse(index model.CommentIndex, content model.CommentContent) model.CommentResponse {
	return model.CommentResponse{
		CommentId:   int64(index.ID),
		ResourceId:  index.ResourceID,
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
