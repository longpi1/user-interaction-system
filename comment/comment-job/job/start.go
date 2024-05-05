package job

import (
	"comment-job/job/comment"
	"comment-job/libary/conf"
	"context"

	"github.com/longpi1/gopkg/libary/queue"
)

func Start() {
	// 注册评论队列的消费者
	comment.RegisterCommentConsumer()
	// 启动队列进行消费
	queue.StartConsumersListener(context.Background(), conf.GetConfig().QueueConfig.Config)
}
