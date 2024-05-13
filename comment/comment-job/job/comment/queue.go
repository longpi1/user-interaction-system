package comment

import (
	"context"
	"encoding/json"

github.com/longpi1/user-interaction-system/comment-job"/libary/conf"

service_queue "github.com/longpi1/user-interaction-system/comment-job/model/service/queue"

"github.com/longpi1/gopkg/libary/queue"
)

var commentQueue = CommentQueue{}

type CommentQueue struct {
}

type CommentInfo struct {
	CommentId     int64  // 评论id
	ResourceId    int64  // 评论所关联的资源id
	ResourceTitle string // 资源的title
	Content       string // 文本信息
	ContentMeta   string // 存储一些关键的附属信息
	ContentRich   string
	Pid           int64  // 父评论 ID
	Ext           string // 额外信息存储
	IP            string
	IPArea        string
	UserID        int64  // 发表者id
	UserName      string // 发表者名称
}

func (queue CommentQueue) GetTopic() string {
	topicName := conf.GetConfig().QueueConfig.TopicName
	return topicName
}

func (queue CommentQueue) Handle(ctx context.Context, msg queue.Msg) (err error) {
	var commentInfo CommentInfo
	if err = json.Unmarshal(msg.Body, &commentInfo); err != nil {
		return err
	}

	return service_queue.UpdateComment(commentInfo)
}
