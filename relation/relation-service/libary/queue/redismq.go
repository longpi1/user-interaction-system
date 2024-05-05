package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"relation-service/model/dao/cache/redis"
	"time"

	"github.com/longpi1/gopkg/libary/log"

	redis2 "github.com/go-redis/redis"
)

type RedisMq struct {
	poolName  string
	groupName string
	timeout   int64
}

type RedisOption struct {
	Timeout int64
}

// SendMsg 按字符串类型生产数据
func (r *RedisMq) SendMsg(topic string, body string) (mqMsg Msg, err error) {
	return r.SendByteMsg(topic, []byte(body))
}

// SendByteMsg 生产数据
func (r *RedisMq) SendByteMsg(topic string, body []byte) (mqMsg Msg, err error) {
	if r.poolName == "" {
		return mqMsg, fmt.Errorf("RedisMq producer not register")
	}

	if topic == "" {
		return mqMsg, fmt.Errorf("RedisMq topic is empty")
	}

	mqMsg = Msg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     getRandMsgId(),
		Body:      body,
		Timestamp: time.Now(),
	}

	data, err := json.Marshal(mqMsg)
	if err != nil {
		return
	}

	key := r.genKey(r.groupName, topic)
	if err = redis.LPush(key, data); err != nil {
		return
	}

	if r.timeout > 0 {
		if err = redis.Expire(key, time.Duration(r.timeout)); err != nil {
			return
		}
	}

	return
}

func (r *RedisMq) SendDelayMsg(topic string, body string, delaySecond int64) (msg Msg, err error) {
	if delaySecond < 1 {
		return r.SendMsg(topic, body)
	}

	if r.poolName == "" {
		err = fmt.Errorf("SendDelayMsg RedisMq not register")
		return
	}

	if topic == "" {
		err = fmt.Errorf("SendDelayMsg RedisMq topic is empty")
		return
	}

	msg = Msg{
		RunType:   SendMsg,
		Topic:     topic,
		MsgId:     getRandMsgId(),
		Body:      []byte(body),
		Timestamp: time.Now(),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return
	}

	var (
		key          = r.genKey(r.groupName, "delay:"+topic)
		expireSecond = time.Now().Unix() + delaySecond
		timePiece    = fmt.Sprintf("%s:%d", key, expireSecond)
		z            = redis2.Z{Score: float64(expireSecond), Member: timePiece}
	)

	if err = redis.ZAdd(key, z); err != nil {
		return
	}

	if err = redis.RPush(timePiece, data); err != nil {
		return
	}

	// consumer will also delete the item
	if r.timeout > 0 {
		_ = redis.Expire(timePiece, time.Duration(r.timeout+delaySecond))
		_ = redis.Expire(key, time.Duration(r.timeout))
	}

	return
}

// ListenReceiveMsgDo 消费数据
func (r *RedisMq) ListenReceiveMsgDo(topic string, receiveDo func(mqMsg Msg)) (err error) {
	if r.poolName == "" {
		return fmt.Errorf("RedisMq producer not register")
	}
	if topic == "" {
		return fmt.Errorf("RedisMq topic is empty")
	}

	var (
		key  = r.genKey(r.groupName, topic)
		key2 = r.genKey(r.groupName, "delay:"+topic)
	)

	go func() {
		for range time.Tick(300 * time.Millisecond) {
			mqMsgList := r.loopReadQueue(key)
			for _, mqMsg := range mqMsgList {
				receiveDo(mqMsg)
			}
		}
	}()

	go func() {
		mqMsgCh, errCh := r.loopReadDelayQueue(key2)
		for mqMsg := range mqMsgCh {
			receiveDo(mqMsg)
		}
		for err = range errCh {
			if err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
				log.Error("ListenReceiveMsgDo Delay topic:%v, err:%+v", topic, err)
			}
		}
	}()

	select {}
}

// 生成队列key
func (r *RedisMq) genKey(groupName string, topic string) string {
	return fmt.Sprintf("comment:%s_%s", groupName, topic)
}

func (r *RedisMq) loopReadQueue(key string) (msgList []Msg) {
	for {
		data, err := redis.Prop(key)
		if err != nil {
			log.Error("loopReadQueue redis RPOP err:%+v", err)
			break
		}

		var msg Msg
		if err = json.Unmarshal([]byte(data), &msg); err != nil {
			log.Error("loopReadQueue Scan err:%+v", err)
			break
		}

		if msg.MsgId != "" {
			msgList = append(msgList, msg)
		}
	}
	return msgList
}

func RegisterRedisMqProducer(connOpt RedisOption, groupName string) (client Producer) {
	return RegisterRedisMq(connOpt, groupName)
}

// RegisterRedisMqConsumer 注册消费者
func RegisterRedisMqConsumer(connOpt RedisOption, groupName string) (client Consumer) {
	return RegisterRedisMq(connOpt, groupName)
}

// RegisterRedisMq 注册redis实例
func RegisterRedisMq(connOpt RedisOption, groupName string) *RedisMq {
	return &RedisMq{
		poolName:  fmt.Sprintf("%s-%d", groupName, time.Now().UnixNano()),
		groupName: groupName,
		timeout:   connOpt.Timeout,
	}
}

func (r *RedisMq) loopReadDelayQueue(key string) (resCh chan Msg, errCh chan error) {
	resCh = make(chan Msg)
	errCh = make(chan error, 1)

	go func() {
		defer close(resCh)
		defer close(errCh)

		for {
			val, err := redis.ZRangeByScore(key, redis2.ZRangeBy{Offset: 0})
			if err != nil {
				return
			}

			for _, listK := range val {
				for {
					pop, err := redis.LPop(listK)
					if err != nil {
						errCh <- err
						return
					} else {
						var msg Msg
						if err = json.Unmarshal([]byte(pop), &msg); err != nil {
							log.Error("loopReadDelayQueue Scan err:%+v", err)
							break
						}

						if msg.MsgId == "" {
							continue
						}

						select {
						case resCh <- msg:
						}
					}
				}
			}
		}
	}()
	return resCh, errCh
}
