package data

import (
	"context"
	"fmt"
	"strconv"

	"github.com/longpi1/user-interaction-system/comment-service/model/dao/cache"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/singleflight"
)

func GetCommentList(param model.CommentParamsList) (model.CommentListResponse, error) {
	key := cache.GetCommentListKey(param.ResourceId, param.Pid)
	// 1.localcache相关操作
	if response, err := cache.GetCommentListFromLocalCache(key); err == nil {
		return response, nil
	}

	// 2.redis相关操作
	if response, err := cache.GetCommentListFromRedisCache(key); err == nil {
		return response, nil
	}

	// 3.db查找查找评论索引集合
	commentIndexs, err := model.GetCommentIndexList(param)
	if err != nil {
		log.Error("数据库获取评论数据失败:%v", err)
		return model.CommentListResponse{}, fmt.Errorf("获取评论数据失败")
	}

	var commentResponses []model.CommentResponse
	for _, index := range commentIndexs {
		// 根据index id 查找对应的content
		if commentContent, err := model.FindCommentContentByCommentId(index.ID); err == nil {
			commentResponse := FormatCommentResponse(index, commentContent)
			commentResponses = append(commentResponses, commentResponse)
		} else {
			log.Error("数据库获取评论内容失败:%v, id: %v", err, index.ID)
		}
	}
	if err != nil {
		log.Error("数据库获取评论数失败:%v, id: %v", err, param.ResourceId)
	}
	commentListResponse := model.CommentListResponse{
		CommentResponses: commentResponses,
	}

	// 4.更新缓存
	cache.SetCommentListToLocalCache(key, commentListResponse)
	cache.SetCommentListToRedisCache(key, commentListResponse)

	return commentListResponse, err
}

// GetCommentListBySingleFlight 基于SingleFlight 获取评论数据
func GetCommentListBySingleFlight(sg *singleflight.Group, ctx context.Context, param model.CommentParamsList) (model.CommentListResponse, error) {
	result, err, _ := sg.Do(strconv.FormatInt(param.ResourceId, 10), func() (interface{}, error) {
		return GetCommentList(param)
	})

	return result.(model.CommentListResponse), err
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
