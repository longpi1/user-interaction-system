package queue

import (
	"context"

	"github.com/longpi1/gopkg/libary/queue"
)

var commentQueue = CommentQueue{}

type CommentQueue struct {
	TopicName string
}

func (queue CommentQueue) GetTopic() string {
	return ""
}

func (queue CommentQueue) Handle(ctx context.Context, msg queue.Msg) (err error) {
	// todo
	return err
}

func Register() {
	queue.RegisterConsumer(commentQueue)
}
