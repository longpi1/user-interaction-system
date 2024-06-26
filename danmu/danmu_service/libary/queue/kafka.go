package queue

import (
	"context"
	"fmt"
	"time"

	"github.com/longpi1/gopkg/libary/log"

	"github.com/IBM/sarama"
)

type Kafka struct {
	Partitions  int32
	producerIns sarama.AsyncProducer
	consumerIns sarama.ConsumerGroup
}

type KafkaConfig struct {
	ClientId    string
	Brokers     []string
	GroupID     string
	Partitions  int32
	Replication int16
	Version     string
	UserName    string
	Password    string
}

// SendMsg 按字符串类型生产数据
func (r *Kafka) SendMsg(topic string, body string) (msg Msg, err error) {
	return r.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 生产数据
func (r *Kafka) SendByteMsg(topic string, body []byte) (msg Msg, err error) {
	producerMessage := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(body),
		Timestamp: time.Now(),
	}

	if r.producerIns == nil {
		err = fmt.Errorf("comment kafka producerIns is nil")
		return
	}

	r.producerIns.Input() <- producerMessage
	sendCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case info := <-r.producerIns.Successes():
		return Msg{
			RunType:   SendMsg,
			Topic:     info.Topic,
			Offset:    info.Offset,
			Partition: info.Partition,
			Timestamp: info.Timestamp,
		}, nil
	case fail := <-r.producerIns.Errors():
		if nil != fail {
			return msg, fail.Err
		}
	case <-sendCtx.Done():
		return msg, fmt.Errorf("send mqMst timeout")
	}
	return msg, nil
}

func (r *Kafka) SendDelayMsg(topic string, body string, delaySecond int64) (msg Msg, err error) {

	return
}

// ListenReceiveMsgDo 消费数据
func (r *Kafka) ListenReceiveMsgDo(topic string, receiveDo func(msg Msg)) (err error) {
	if r.consumerIns == nil {
		return fmt.Errorf("comment kafka consumer not register")
	}

	consumer := KaConsumer{
		ready:        make(chan bool),
		receiveDoFun: receiveDo,
	}

	consumerCtx, cancel := context.WithCancel(context.Background())
	go func(consumerCtx context.Context) {
		for {
			if err = r.consumerIns.Consume(consumerCtx, []string{topic}, &consumer); err != nil {
				log.Error("kafka Error from consumer, err%+v", err)
			}

			if consumerCtx.Err() != nil {
				log.Error(fmt.Sprintf("kafka consoumer stop : %v", consumerCtx.Err()))
				return
			}
			consumer.ready = make(chan bool)
		}
	}(consumerCtx)

	// await till the consumer has been set up
	<-consumer.ready
	log.Debug("kafka consumer up and running!...")

	func(args ...interface{}) {
		log.Debug("kafka consumer close...")
		cancel()
		if err = r.consumerIns.Close(); err != nil {
			log.Error("kafka Error closing client, err:%+v", err)
		}
	}()
	return
}

// RegisterKafkaConsumer 注册消费者
func RegisterKafkaConsumer(connOpt KafkaConfig) (client Consumer, err error) {
	mqIns := &Kafka{}
	kfkVersion, err := sarama.ParseKafkaVersion(connOpt.Version)
	if err != nil {
		return
	}
	if !validateVersion(kfkVersion) {
		kfkVersion = sarama.V2_4_0_0
	}

	brokers := connOpt.Brokers
	conf := sarama.NewConfig()
	conf.Consumer.Return.Errors = true
	conf.Version = kfkVersion
	if connOpt.UserName != "" {
		conf.Net.SASL.Enable = true
		conf.Net.SASL.User = connOpt.UserName
		conf.Net.SASL.Password = connOpt.Password
	}

	// 默认按随机方式消费
	conf.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	conf.Consumer.Offsets.AutoCommit.Interval = 10 * time.Millisecond
	conf.ClientID = connOpt.ClientId

	consumerClient, err := sarama.NewConsumerGroup(brokers, connOpt.GroupID, conf)
	if err != nil {
		return
	}
	mqIns.consumerIns = consumerClient
	return mqIns, err
}

// RegisterKafkaProducer 注册并启动生产者接口实现
func RegisterKafkaProducer(connOpt KafkaConfig) (client Producer, err error) {
	mqIns := &Kafka{}
	connOpt.ClientId = "producer"

	// 这里如果使用go程需要处理chan同步问题
	if err = doRegisterKafkaProducer(connOpt, mqIns); err != nil {
		return nil, err
	}

	return mqIns, nil
}

// doRegisterKafkaProducer 注册同步类型实例
func doRegisterKafkaProducer(connOpt KafkaConfig, mqIns *Kafka) (err error) {
	kfkVersion, err := sarama.ParseKafkaVersion(connOpt.Version)
	if err != nil {
		return
	}
	if !validateVersion(kfkVersion) {
		kfkVersion = sarama.V2_4_0_0
	}

	brokers := connOpt.Brokers
	conf := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	conf.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向partition发送消息
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	conf.Producer.Return.Successes = true

	conf.Producer.Return.Errors = true
	conf.Producer.Compression = sarama.CompressionNone
	conf.ClientID = connOpt.ClientId

	conf.Version = kfkVersion
	if connOpt.UserName != "" {
		conf.Net.SASL.Enable = true
		conf.Net.SASL.User = connOpt.UserName
		conf.Net.SASL.Password = connOpt.Password
	}

	mqIns.producerIns, err = sarama.NewAsyncProducer(brokers, conf)
	if err != nil {
		return
	}

	func(args ...interface{}) {
		log.Info("kafka producer AsyncClose...")
		mqIns.producerIns.AsyncClose()
	}()
	return
}

// validateVersion 验证版本是否有效
func validateVersion(version sarama.KafkaVersion) bool {
	for _, item := range sarama.SupportedVersions {
		if version.String() == item.String() {
			return true
		}
	}
	return false
}

type KaConsumer struct {
	ready        chan bool
	receiveDoFun func(msg Msg)
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *KaConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *KaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *KaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	// `ConsumeClaim` 方法已经是 goroutine 调用 不要在该方法内进行 goroutine
	for message := range claim.Messages() {
		consumer.receiveDoFun(Msg{
			RunType:   ReceiveMsg,
			Topic:     message.Topic,
			Body:      message.Value,
			Offset:    message.Offset,
			Timestamp: message.Timestamp,
			Partition: message.Partition,
		})
		session.MarkMessage(message, "")
	}
	return nil
}
