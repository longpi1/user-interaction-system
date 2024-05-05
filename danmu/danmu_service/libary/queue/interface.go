package queue

import (
	"fmt"
	"math/rand"
	"relation-service/libary/constant"
	"relation-service/model/dao/cache/redis"
	"strconv"
	"sync"
	"time"
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
	Pulsar    PulsarConf
}

type RedisConf struct {
	Timeout int64 `json:"timeout"`
}
type RocketConf struct {
	Address  []string `json:"address"`
	LogLevel string   `json:"logLevel"`
}

type PulsarConf struct {
	Address          []string `json:"address"`
	LogLevel         string   `json:"logLevel"`
	Topic            string   `json:"topic"`
	URL              string   `json:"url"`
	Type             int      `json:"type"`
	SubscriptionName string   `json:"subscriptionName"`
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
func InstanceProducer() (client Producer, err error) {
	return NewProducer(config.GroupName)
}

// NewProducer 初始化生产者实例
func NewProducer(groupName string) (client Producer, err error) {
	if groupName == "" {
		err = fmt.Errorf("mq groupName is empty")
		return
	}

	switch config.Driver {
	case constant.RocketMqName:
		if len(config.Rocket.Address) == 0 {
			err = fmt.Errorf("comment rocketmq address is not support")
			return
		}
		client, err = RegisterRocketProducer(config.Rocket.Address, groupName, config.Retry)
	case constant.KafkaMqName:
		if len(config.Kafka.Address) == 0 {
			err = fmt.Errorf("comment kafka address is not support")
			return
		}
		client, err = RegisterKafkaProducer(KafkaConfig{
			Brokers: config.Kafka.Address,
			GroupID: groupName,
			Version: config.Kafka.Version,
		})
	case constant.RedisMqName:
		if _, err = redis.Ping(); err == nil {
			client = RegisterRedisMq(RedisOption{
				Timeout: config.Redis.Timeout,
			}, groupName)
		}
	case constant.PulsarMqName:
		if len(config.Pulsar.Address) == 0 {
			err = fmt.Errorf("comment pulsar address is not support")
			return
		}
		client, err = RegisterPulsarProducer(config.Pulsar)
	default:
		err = fmt.Errorf("comment driver is not support")
	}

	if err != nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	return
}

// NewConsumer 初始化消费者实例
func NewConsumer(groupName string) (client Consumer, err error) {
	if groupName == "" {
		err = fmt.Errorf("mq groupName is empty")
		return
	}

	switch config.Driver {
	case constant.RocketMqName:
		if len(config.Rocket.Address) == 0 {
			err = fmt.Errorf("comment.rocketmq.address is empty")
			return
		}
		client, err = RegisterRocketConsumer(config.Rocket.Address, groupName)
	case constant.KafkaMqName:
		if len(config.Kafka.Address) == 0 {
			err = fmt.Errorf("comment kafka address is not support")
			return
		}

		randTag := strconv.FormatInt(time.Now().Unix(), 10)
		// 是否支持创建多个消费者
		if !config.Kafka.MultiConsumer {
			randTag = "001"
		}

		clientId := "HOTGO-Consumer-" + groupName
		if config.Kafka.RandClient {
			clientId += "-" + randTag
		}

		client, err = RegisterKafkaConsumer(KafkaConfig{
			Brokers:  config.Kafka.Address,
			GroupID:  groupName,
			Version:  config.Kafka.Version,
			ClientId: clientId,
		})
	case constant.RedisMqName:
		if _, err = redis.Ping(); err == nil {
			client = RegisterRedisMqConsumer(RedisOption{
				Timeout: config.Redis.Timeout,
			}, groupName)
		}
	case constant.PulsarMqName:
		if len(config.Pulsar.Address) == 0 {
			err = fmt.Errorf("comment pulsar address is not support")
			return
		}
		client, err = RegisterPulsarConsumer(config.Pulsar)
	default:
		err = fmt.Errorf("comment driver is not support")
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

func getRandMsgId() string {
	rand.NewSource(time.Now().UnixNano())
	radium := rand.Intn(999) + 1
	timeCode := time.Now().UnixNano()
	return fmt.Sprintf("%d%.4d", timeCode, radium)
}
