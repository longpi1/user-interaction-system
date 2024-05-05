package comment

import (
	"comment-job/libary/conf"
	"context"

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
	// todo
	return err
}
