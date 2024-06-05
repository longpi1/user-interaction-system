package relation

import (
	"context"
	"encoding/json"

	"github.com/longpi1/user-interaction-system/relation-job/libary/conf"
	"github.com/longpi1/user-interaction-system/relation-service/model/dao/db/model"

	service_queue "github.com/longpi1/user-interaction-system/relation-job/model/service/queue"

	"github.com/longpi1/gopkg/libary/queue"
)

var relationQueue = RelationQueue{}

type RelationQueue struct {
}

func (queue RelationQueue) GetTopic() string {
	topicName := conf.GetConfig().QueueConfig.TopicName
	return topicName
}

func (queue RelationQueue) Handle(ctx context.Context, msg queue.Msg) (err error) {
	var relation model.Relation
	if err = json.Unmarshal(msg.Body, &relation); err != nil {
		return err
	}

	return service_queue.Relation(relation)
}
