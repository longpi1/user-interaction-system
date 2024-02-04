package cache

import (
	"github.com/go-redis/redis"
	"sync"
	"user-interaction-system/libary/conf"
	"user-interaction-system/libary/log"
)

var redisClient *redis.Client
var once sync.Once

func GetClient() (*redis.Client, error) {
	var err error
	if redisClient == nil {
		once.Do(func() {
			redisClient, err = NewClient(conf.RedisConf{})
			if err != nil {
				log.Fatal("cache run err", err)
			}
		})
	}
	// 测试连接是否成功
	err = redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

func NewClient(config conf.RedisConf) (*redis.Client, error) {
	// 创建Redis客户端对象
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Network,
		Password: config.Password, // 如果有密码，请填写密码
		DB:       config.DB,       // 选择要使用的数据库编号
	})
	// 测试连接是否成功
	err := redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}
