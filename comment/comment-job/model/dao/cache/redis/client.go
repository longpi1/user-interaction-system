package redis

import (
	"sync"

	"comment-job/libary/conf"

	"github.com/longpi1/gopkg/libary/log"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client
var once sync.Once

func GetClient() *redis.Client {
	var err error
	if redisClient == nil {
		once.Do(func() {
			redisClient, err = NewClient(conf.GetConfig())
			if err != nil {
				log.Fatal("redis cache run err", err)
			}
		})
	}

	return redisClient
}

func NewClient(config *conf.WebConfig) (*redis.Client, error) {
	// 创建Redis客户端对象
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Address,
		Password: config.RedisConfig.Password, // 如果有密码，请填写密码
		DB:       config.RedisConfig.Db,       // 选择要使用的数据库编号
	})
	// 测试连接是否成功
	err := redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}
