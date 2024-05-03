package event

import (
	"comment-job/libary/constant"
	"comment-job/libary/event"
	"comment-job/libary/msg"
	"comment-job/model/dao/db/model"
	"reflect"

	"github.com/longpi1/gopkg/libary/log"
)

func init() {
	event.RegisterHandler(reflect.TypeOf(event.CommentUpdateEvent{}), addComment)
}

func addComment(comment interface{}) {
	e := comment.(event.CommentUpdateEvent)

	commentIndex, err := model.FindCommentIndexById(e.CommentId)
	if err != nil {
		log.Error("数据库评论查询失败，%v", comment)
	}
	// 发送消息
	handleMsg(commentIndex)
}

// 处理评论消息
func handleMsg(comment *model.CommentIndex) {
	commentMsg := getCommentMsg(comment)
	if commentMsg == nil {
		return
	}
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

	// todo 消息中心发送数据

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
	if commentMsg.ParentComment != nil && commentMsg.ParentComment.UserID == to {
		return
	}

	//repliedContent, err := model.FindCommentContentByCommentId(queue.ID)
	//if err != nil {
	//	log.Error("数据库获取评论内容失败： %v", err)
	//}

	// todo 消息中心发送数据

}

// handleQuoteMsg 给被引用人发送消息
func handleQuoteMsg(comment *model.CommentIndex, commentMsg *CommentMsg) {
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
	//repliedContent, err := model.FindCommentContentByCommentId(queue.ID)
	//if err != nil {
	//	log.Error("数据库获取评论内容失败： %v", err)
	//}

	// todo 消息中心发送数据
}

func getCommentMsg(comment *model.CommentIndex) *CommentMsg {
	if comment.State != constant.StateReviewed {
		log.Error("未过审核暂不需处理")
		return nil
	}
	switch comment.ResourceType {
	// 如果是回复时赋值父评论消息
	case constant.ResourceComment:
		parentComment, err := model.FindCommentIndexById(comment.PID)
		if err != nil || parentComment == nil {
			log.Error("数据库获取父评论失败: %v", err)
			return nil
		}
		if parentComment == nil || parentComment.State != constant.StateReviewed {
			log.Error("父评论未过审核暂不需处理")
			return nil
		}
		ret := &CommentMsg{
			Comment:       comment,
			ResourceType:  parentComment.ResourceType, // 二级评论时，取一级评论的
			ResourceID:    parentComment.ResourceID,   // 二级评论时，取一级评论的
			ParentComment: parentComment,              // 一级评论
		}
		return ret
	default:
		ret := &CommentMsg{
			Comment:       comment,
			ResourceType:  comment.ResourceType,
			ResourceID:    comment.ResourceID,
			ParentComment: nil,
		}
		return ret
	}
}

type CommentMsg struct {
	ResourceType  string              // 资源类型
	ResourceID    int64               // 资源ID
	Comment       *model.CommentIndex // 当前评论
	ParentComment *model.CommentIndex // 上一级评论（二级评论的时候有值）
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
func (c *CommentMsg) msgContent(commentID int) string {
	commentContent, err := model.FindCommentContentByCommentId(int64(commentID))
	if err != nil {
		return ""
	}
	return commentContent.Content
}

// msgRepliedContent 被回复的内容
func (c *CommentMsg) msgRepliedContent() string {

	return ""
}

func (c *CommentMsg) rootEntityUserId() int64 {
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
