package queue

import (
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type Queue interface {
	Start() error
	Stop()
}

type Producer interface {
	SendMsg(topic string, body string) (msg Msg, err error)
	SendByteMsg(topic string, body []byte) (msg Msg, err error)
	SendDelayMsg(topic string, body string, delaySecond int64) (mqMsg Msg, err error)
}

type Consumer interface {
	ListenReceiveMsgDo(topic string, receiveDo func(Msg Msg)) (err error)
}

const (
	_ = iota
	SendMsg
	ReceiveMsg
)

type Config struct {
	Switch    bool   `json:"switch"`
	Driver    string `json:"driver"`
	Retry     int    `json:"retry"`
	GroupName string `json:"groupName"`
	Redis     RedisConf
	Rocket    RocketConf
	Kafka     KafkaConf
}

type RedisConf struct {
	Timeout int64 `json:"timeout"`
}
type RocketConf struct {
	Address  []string `json:"address"`
	LogLevel string   `json:"logLevel"`
}

type KafkaConf struct {
	Address       []string `json:"address"`
	Version       string   `json:"version"`
	RandClient    bool     `json:"randClient"`
	MultiConsumer bool     `json:"multiConsumer"`
}

type Msg struct {
	RunType   int       `json:"run_type"`
	Topic     string    `json:"topic"`
	MsgId     string    `json:"msg_id"`
	Offset    int64     `json:"offset"`
	Partition int32     `json:"partition"`
	Timestamp time.Time `json:"timestamp"`
	Body      []byte    `json:"body"`
}

var (
	mutex  sync.Mutex
	config Config
)

// InstanceConsumer 实例化消费者
func InstanceConsumer() (mqClient Consumer, err error) {
	return NewConsumer(config.GroupName)
}

// InstanceProducer 实例化生产者
func InstanceProducer() (mqClient Producer, err error) {
	return NewProducer(config.GroupName)
}

// NewProducer 初始化生产者实例
func NewProducer(groupName string) (mqClient Producer, err error) {

	if groupName == "" {
		err = fmt.Errorf("mq groupName is empty.")
		return
	}

	switch config.Driver {
	case "rocketmq":
		if len(config.Rocketmq.Address) == 0 {
			err = fmt.Errorf("queue rocketmq address is not support")
			return
		}
		mqClient, err = RegisterRocketProducer(config.Rocketmq.Address, groupName, config.Retry)
	case "kafka":
		if len(config.Kafka.Address) == 0 {
			err = fmt.Errorf("queue kafka address is not support")
			return
		}
		mqClient, err = RegisterKafkaProducer(KafkaConfig{
			Brokers: config.Kafka.Address,
			GroupID: groupName,
			Version: config.Kafka.Version,
		})
	case "redis":
		if _, err = g.Redis().Do(ctx, "ping"); err == nil {
			mqClient = RegisterRedisProducer(RedisOption{
				Timeout: config.Redis.Timeout,
			}, groupName)
		}
	default:
		err = fmt.Errorf("queue driver is not support")
	}

	if err != nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	return
}

// NewConsumer 初始化消费者实例
func NewConsumer(groupName string) (mqClient Consumer, err error) {
	if groupName == "" {
		err = fmt.Errorf("mq groupName is empty.")
		return
	}

	switch config.Driver {
	case "rocketmq":
		if len(config.Rocketmq.Address) == 0 {
			err = fmt.Errorf("queue.rocketmq.address is empty.")
			return
		}
		mqClient, err = RegisterRocketConsumer(config.Rocketmq.Address, groupName)
	case "kafka":
		if len(config.Kafka.Address) == 0 {
			err = fmt.Errorf("queue kafka address is not support")
			return
		}

		randTag := string(charset.RandomCreateBytes(6))
		// 是否支持创建多个消费者
		if !config.Kafka.MultiConsumer {
			randTag = "001"
		}

		clientId := "HOTGO-Consumer-" + groupName
		if config.Kafka.RandClient {
			clientId += "-" + randTag
		}

		mqClient, err = RegisterKafkaConsumer(KafkaConfig{
			Brokers:  config.Kafka.Address,
			GroupID:  groupName,
			Version:  config.Kafka.Version,
			ClientId: clientId,
		})
	case "redis":
		if _, err = g.Redis().Do(ctx, "ping"); err == nil {
			mqClient = RegisterRedisConsumer(RedisOption{
				Timeout: config.Redis.Timeout,
			}, groupName)
		}
	case "disk":
		config.Disk.GroupName = groupName
		mqClient, err = RegisterDiskConsumer(config.Disk)
	default:
		err = fmt.Errorf("queue driver is not support")
	}

	if err != nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	return
}

// BodyString 返回消息体
func (m *Msg) BodyString() string {
	return string(m.Body)
}
