package comment

import "github.com/longpi1/gopkg/libary/queue"

func RegisterCommentConsumer() {
	queue.RegisterConsumer(commentQueue)
}
