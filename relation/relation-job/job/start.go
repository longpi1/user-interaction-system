package job

import (
	"context"

	"github.com/longpi1/user-interaction-system/relation-job/libary/conf"

	"github.com/longpi1/gopkg/libary/queue"
)

func Start() {
	// 注册评论队列的消费者

	// 启动队列进行消费
	queue.StartConsumersListener(context.Background(), conf.GetConfig().QueueConfig.Config)
}
