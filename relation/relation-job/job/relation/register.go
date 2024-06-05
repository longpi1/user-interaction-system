package relation

import "github.com/longpi1/gopkg/libary/queue"

func RegisterRelationConsumer() {
	queue.RegisterConsumer(relationQueue)
}
