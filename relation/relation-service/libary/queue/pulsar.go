package queue

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Pulsar struct {
	Client   pulsar.Client
	Producer pulsar.Producer
	Consumer pulsar.Consumer
}

type PulsarConfig struct {
	ClientId    string
	Brokers     []string
	GroupID     string
	Partitions  int32
	Replication int16
	Version     string
	UserName    string
	Password    string
}

// NewPulsar creates a new client with the given service URL.
func NewPulsar(serviceURL string) (*Pulsar, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: serviceURL,
	})
	if err != nil {
		return nil, fmt.Errorf("could not create pulsar client: %v", err)
	}

	return &Pulsar{Client: client}, nil
}

// RegisterPulsarConsumer creates a consumer for a specific topic and subscription.
func (p *Pulsar) RegisterPulsarConsumer(topic, subscriptionName string) error {
	consumer, err := p.Client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subscriptionName,
		Type:             pulsar.Shared,
	})
	if err != nil {
		return fmt.Errorf("could not create consumer: %v", err)
	}
	p.Consumer = consumer
	return nil
}

// RegisterPulsarProducer creates a producer for a specific topic.
func (p *Pulsar) RegisterPulsarProducer(topic string) error {
	producer, err := p.Client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		return fmt.Errorf("could not create producer: %v", err)
	}
	p.Producer = producer
	return nil
}

// SendMsg 按字符串类型生产数据
func (p *Pulsar) SendMsg(topic string, body string) (msg Msg, err error) {
	return p.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 生产数据
func (p *Pulsar) SendByteMsg(topic string, body []byte) (msg Msg, err error) {
	if p.Producer == nil {
		return msg, fmt.Errorf("producer is not set")
	}

	messageID, err := p.Producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: body,
	})
	if err != nil {
		return msg, fmt.Errorf("could not send message: %d, %v", messageID, err)
	}

	msg = Msg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     messageID.String(),
		Body:      body,
		Timestamp: time.Now(),
	}

	return msg, err
}

func (p *Pulsar) SendDelayMsg(topic string, body string, delaySecond int64) (msg Msg, err error) {

	return
}

// ListenReceiveMsgDo 消费数据
func (p *Pulsar) ListenReceiveMsgDo(topic string, receiveDo func(msg Msg)) (err error) {
	if p.Consumer == nil {
		return fmt.Errorf("consumer is not set")
	}
	go func() {
		for {
			data, err := p.Consumer.Receive(context.Background())
			if err != nil {
				log.Printf("Error receiving message: %v", err)
				continue
			}
			msg := Msg{
				RunType:   SendMsg,
				Topic:     topic,
				MsgId:     getRandMsgId(),
				Body:      data.Payload(),
				Timestamp: time.Now(),
			}

			receiveDo(msg)
			if err != nil {
				log.Printf("Error handling message: %v", err)
				// Consider what to do with the message: Ack/Nack
				p.Consumer.Nack(data)
			} else {
				p.Consumer.Ack(data)
			}
		}
	}()

	return nil
}

// Close closes the client and releases all resources.
func (p *Pulsar) Close() {
	if p.Producer != nil {
		p.Producer.Close()
	}
	if p.Consumer != nil {
		p.Consumer.Close()
	}
	p.Client.Close()
}
