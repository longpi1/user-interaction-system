package queue

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/gogf/gf/v2/frame/g"
)

type RocketMq struct {
	endPoints   []string
	producerIns rocketmq.Producer
	consumerIns rocketmq.PushConsumer
}

// rewriteLog 重写日志
func rewriteLog() {
	rlog.SetLogger(&RocketMqLogger{Flag: "[rocket_mq]", LevelLog: g.Cfg().MustGet(ctx, "queue.rocketmq.logLevel", "debug").String()})
}

// RegisterRocketProducer 注册并启动生产者接口实现
func RegisterRocketProducer(endPoints []string, groupName string, retry int) (client MqProducer, err error) {
	rewriteLog()
	client, err = RegisterRocketMqProducer(endPoints, groupName, retry)
	if err != nil {
		return
	}
	return
}

// RegisterRocketConsumer 注册消费者
func RegisterRocketConsumer(endPoints []string, groupName string) (client MqConsumer, err error) {
	rewriteLog()
	client, err = RegisterRocketMqConsumer(endPoints, groupName)
	if err != nil {
		return
	}
	return
}

// SendMsg 按字符串类型生产数据
func (r *RocketMq) SendMsg(topic string, body string) (mqMsg msg, err error) {
	return r.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 生产数据
func (r *RocketMq) SendByteMsg(topic string, body []byte) (mqMsg msg, err error) {
	if r.producerIns == nil {
		return mqMsg, fmt.Errorf("rocketMq producer not register")
	}

	result, err := r.producerIns.SendSync(context.Background(), &primitive.Message{
		Topic: topic,
		Body:  body,
	})

	if err != nil {
		return
	}
	if result.Status != primitive.SendOK {
		return mqMsg, fmt.Errorff("rocketMq producer send msg error status:%v", result.Status)
	}

	mqMsg = msg{
		RunType: SendMsg,
		Topic:   topic,
		MsgId:   result.MsgID,
		Body:    body,
	}
	return mqMsg, nil
}

func (r *RocketMq) SendDelayMsg(topic string, body string, delaySecond int64) (mqMsg msg, err error) {
	err = fmt.Errorf("implement me")
	return
}

// ListenReceiveMsgDo 消费数据
func (r *RocketMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg msg)) (err error) {
	if r.consumerIns == nil {
		return fmt.Errorf("rocketMq consumer not register")
	}

	err = r.consumerIns.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, item := range msgs {
			go receiveDo(msg{
				RunType: ReceiveMsg,
				Topic:   item.Topic,
				MsgId:   item.MsgId,
				Body:    item.Body,
			})
		}
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		return
	}

	if err = r.consumerIns.Start(); err != nil {
		_ = r.consumerIns.Unsubscribe(topic)
		return
	}
	return
}

// RegisterRocketMqProducer 注册rocketmq生产者
func RegisterRocketMqProducer(endPoints []string, groupName string, retry int) (mqIns *RocketMq, err error) {
	addr, err := primitive.NewNamesrvAddr(endPoints...)
	if err != nil {
		return nil, err
	}
	mqIns = &RocketMq{
		endPoints: endPoints,
	}

	if retry <= 0 {
		retry = 0
	}

	mqIns.producerIns, err = rocketmq.NewProducer(
		producer.WithNameServer(addr),
		producer.WithRetry(retry),
		producer.WithGroupName(groupName),
	)

	if err != nil {
		return nil, err
	}

	err = mqIns.producerIns.Start()
	if err != nil {
		return nil, err
	}
	return mqIns, nil
}

// RegisterRocketMqConsumer 注册rocketmq消费者
func RegisterRocketMqConsumer(endPoints []string, groupName string) (mqIns *RocketMq, err error) {
	addr, err := primitive.NewNamesrvAddr(endPoints...)
	if err != nil {
		return nil, err
	}
	mqIns = &RocketMq{
		endPoints: endPoints,
	}
	mqIns.consumerIns, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer(addr),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(groupName),
	)

	if err != nil {
		return nil, err
	}
	return mqIns, nil
}
