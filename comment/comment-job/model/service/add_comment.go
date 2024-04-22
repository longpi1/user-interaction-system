package service

import (
	"comment-job/libary/constant"
	"comment-job/libary/event"
	"comment-job/libary/log"
	"comment-job/libary/msg"
	"comment-job/model/dao/db/model"
	"reflect"

	"github.com/spf13/cast"
)

func init() {
	event.RegisterHandler(reflect.TypeOf(event.CommentCreateEvent{}), addComment)
}

func addComment(comment interface{}) {
	e := comment.(event.CommentCreateEvent)

	commentIndex, err := model.FindCommentIndexById(e.CommentId)
	if err != nil {
		log.Error("数据库评论查询失败，%v", comment)
	}
	// 发送消息
	handleMsg(&commentIndex)
}

// 处理评论消息
func handleMsg(comment *model.CommentIndex) {
	commentMsg := getCommentMsg(comment)

	// 给被回复的资源对象作者发送消息
	handleResourceMsgToAuthor(comment, commentMsg)
	// 给被引用人发送消息
	handleQuoteMsg(comment, commentMsg)
	// 处理该评论信息的相关事件
	handleReplyMsg(comment, commentMsg)
}

// handleResourceMsg 给被回复的资源对象作者发送消息
func handleResourceMsgToAuthor(comment *model.CommentIndex, commentMsg *CommentMsg) {
	var (
		from = comment.UserID
		to   = commentMsg.rootEntityUserId()
	)
	if from == to {
		return
	}
	if to <= 0 {
		log.Error("消息发送失败, to: %v", to)
		return
	}

	// 如果回复的评论作者就是帖子作者，那么只给回复作者发消息即可，这里就不再给帖子作者发消息了
	if commentMsg.ParentComment != nil && commentMsg.ParentComment.UserID == to {
		return
	}
	// 同上
	if commentMsg.QuoteComment != nil && commentMsg.QuoteComment.UserID == to {
		return
	}

	services.MessageService.SendMsg(from, to,
		commentMsg.getMsgType(),
		commentMsg.getMsgTitle(),
		commentMsg.msgContent(),
		commentMsg.msgRepliedContent(),
		&msg.CommentExtraData{
			ResourceType: comment.ResourceType,
			ResourceID:   comment.ResourceID,
			PID:          comment.PID,
			RootType:     commentMsg.rootResourceType(),
			RootId:       cast.ToString(commentMsg.rootEntityId()),
		})
}

// handleReplyMsg 处理该评论信息的相关事件
func handleReplyMsg(comment *model.CommentIndex, commentMsg *CommentMsg) {
	if commentMsg.ParentComment == nil {
		return
	}

	var (
		from = comment.UserID
		to   = commentMsg.ParentComment.UserID
	)

	if from == to {
		return
	}

	// 如果回复的评论作者就是被引用消息的作者作者，那么只发引用消息即可
	if commentMsg.QuoteComment != nil && commentMsg.QuoteComment.UserID == to {
		return
	}

	var (
		title   = commentMsg.getMsgTitle()
		content = commentMsg.msgContent()
	)
	repliedContent, err := model.FindCommentContentByCommentId(comment.ID)
	if err != nil {
		log.Error("数据库获取评论内容失败： %v", err)
	}

	services.MessageService.SendMsg(from, to, msg.TypeCommentReply, title, content, repliedContent,
		&msg.CommentExtraData{
			ResourceType: comment.ResourceType,
			ResourceID:   comment.ResourceID,
			PID:          comment.PID,
			RootType:     commentMsg.rootResourceType(),
			RootId:       cast.ToString(commentMsg.rootEntityId()),
		})
}

// handleQuoteMsg 给被引用人发送消息
func handleQuoteMsg(comment *model.CommentIndex, commentMsg *CommentMsg) {
	if commentMsg.QuoteComment == nil {
		return
	}

	var (
		from    = comment.UserID
		to      = commentMsg.QuoteComment.UserID
		title   = commentMsg.getMsgTitle()
		content = commentMsg.msgContent()
	)
	repliedContent, err := model.FindCommentContentByCommentId(comment.ID)
	if err != nil {
		log.Error("数据库获取评论内容失败： %v", err)
	}

	if from == to {
		return
	}

	services.MessageService.SendMsg(from, to, msg.TypeCommentReply, title, content, repliedContent,
		&msg.CommentExtraData{
			ResourceType: comment.ResourceType,
			ResourceID:   comment.ResourceID,
			PID:          comment.PID,
			RootType:     commentMsg.rootResourceType(),
			RootId:       cast.ToString(commentMsg.rootEntityId()),
		})
}

func getCommentMsg(comment *model.CommentIndex) *CommentMsg {
	switch comment.ResourceType {
	case constant.ResourceTopic:
		topic := services.TopicService.Get(comment.EntityId)
		if topic != nil && topic.Status == constant.StatusOk {
			return &CommentMsg{
				Comment:      comment,
				ResourceType: comment.ResourceType,
				ResourceID:   comment.ResourceID,
				Entity:       topic,
			}
		}
	case constant.ResourceArticle:
		article := services.ArticleService.Get(comment.EntityId)
		if article != nil && article.Status == constant.StatusOk {
			return &CommentMsg{
				Comment:      comment,
				ResourceType: comment.ResourceType,
				ResourceID:   comment.ResourceID,
				Entity:       article,
			}
		}
	case constant.ResourceNews:
		return "回复了你的新闻"
	case constant.ResourceVideo:
		return "回复了你的视频"
	case constant.ResourceComment:
		parentComment := services.CommentService.Get(comment.EntityId)
		if parentComment == nil || parentComment.Status != constant.StatusOk {
			return nil
		}

		ret := &CommentMsg{
			Comment:       comment,
			ResourceType:  parentComment.ResourceType, // 二级评论时，取一级评论的
			ResourceID:    parentComment.ResourceID,   // 二级评论时，取一级评论的
			ParentComment: parentComment,              // 一级评论
		}

		if parentComment.ResourceType == constant.ResourceTopic {
			topic := services.TopicService.Get(parentComment.EntityId)
			if topic != nil && topic.Status == constant.StatusOk {
				ret.Entity = topic
			}
		} else if parentComment.ResourceType == constant.ResourceArticle {
			article := services.ArticleService.Get(parentComment.EntityId)
			if article != nil && article.Status == constant.StatusOk {
				ret.Entity = article
			}
		} else {
			return nil
		}

		if comment.QuoteId > 0 { // 三级评论
			quoteComment := services.CommentService.Get(comment.QuoteId)
			if quoteComment != nil && quoteComment.Status == constant.StatusOk {
				ret.QuoteComment = quoteComment
			}
		}

		return ret
	}
	return nil
}

type CommentMsg struct {
	ResourceType  string              // 资源类型
	ResourceID    int64               // 资源ID
	Comment       *model.CommentIndex // 当前评论
	ParentComment *model.CommentIndex // 上一级评论（二级评论的时候有值）
	QuoteComment  *model.CommentIndex // 引用评论
}

// getMsgType 消息类型
func (c *CommentMsg) getMsgType() msg.Type {
	if c.ResourceType == constant.ResourceComment {
		return msg.TypeCommentReply
	}
	return msg.TypeResourceComment
}

// getMsgTitle 消息标题
func (c *CommentMsg) getMsgTitle() string {
	switch c.ResourceType {
	case constant.ResourceTopic:
		return "回复了你的话题"
	case constant.ResourceArticle:
		return "回复了你的文章"
	case constant.ResourceNews:
		return "回复了你的新闻"
	case constant.ResourceVideo:
		return "回复了你的视频"
	case constant.ResourceComment:
		return "回复了你的评论"
	default:
		return ""
	}
}

// msgContent 回复内容
func (c *CommentMsg) msgContent() string {
	return common.GetSummary(c.Comment.ContentType, c.Comment.Content)
}

// msgRepliedContent 被回复的内容
func (c *CommentMsg) msgRepliedContent() string {
	if c.ResourceType == constant.ResourceArticle {
		article := c.Entity.(*models.Article)
		return "《" + article.Title + "》"
	} else if c.ResourceType == constant.ResourceTopic {
		topic := c.Entity.(*models.Topic)
		return "《" + topic.GetTitle() + "》"
	}
	return ""
}

func (c *CommentMsg) rootEntityUserId() int64 {
	if c.ParentComment != nil { // 二级评论
		if c.ParentComment.ResourceType == constant.ResourceTopic {
			topic := c.Entity.(*models.Topic)
			return topic.UserId
		} else if c.ParentComment.ResourceType == constant.ResourceArticle {
			article := c.Entity.(*models.Article)
			return article.UserId
		}
	} else {
		if c.Comment.ResourceType == constant.ResourceTopic {
			topic := c.Entity.(*models.Topic)
			return topic.UserId
		} else if c.Comment.ResourceType == constant.ResourceArticle {
			article := c.Entity.(*models.Article)
			return article.UserId
		}
	}
	return 0
}

func (c *CommentMsg) rootResourceType() string {
	if c.ParentComment != nil { // 二级评论
		return c.ParentComment.ResourceType
	} else {
		return c.Comment.ResourceType
	}
}

func (c *CommentMsg) rootEntityId() int64 {
	if c.ParentComment != nil { // 二级评论
		return c.ParentComment.ResourceID
	} else {
		return c.Comment.ResourceID
	}
}
