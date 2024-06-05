package event

import (
	"reflect"

	"github.com/longpi1/user-interaction-system/relation-job/libary/event"
)

func init() {
	event.RegisterHandler(reflect.TypeOf(event.RelationUpdateEvent{}), relation)
}

func relation(relation interface{}) {
	relationEvent := relation.(event.RelationUpdateEvent)

	// 发送消息
	handleMsg(relationEvent)
}

// 处理消息
func handleMsg(relationEvent event.RelationUpdateEvent) {

}
