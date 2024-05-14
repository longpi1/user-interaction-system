package comment

import (
	"context"
	"encoding/json"

	"github.com/longpi1/user-interaction-system/comment-job/libary/conf"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"

	service_queue "github.com/longpi1/user-interaction-system/comment-job/model/service/queue"

	"github.com/longpi1/gopkg/libary/queue"
)

var commentQueue = CommentQueue{}

type CommentQueue struct {
}

func (queue CommentQueue) GetTopic() string {
	topicName := conf.GetConfig().QueueConfig.TopicName
	return topicName
}

func (queue CommentQueue) Handle(ctx context.Context, msg queue.Msg) (err error) {
	var commentInfo model.CommentInfo
	if err = json.Unmarshal(msg.Body, &commentInfo); err != nil {
		return err
	}

	return service_queue.UpdateComment(commentInfo)
}
