package redis

import (
	"github.com/go-redis/redis"
	"sync"
	"user-interaction-system/libary/conf"
)

var redisClient *redis.Client
var once sync.Once

func GetClient() *redis.Client {
	return redisClient
}

func NewClient(config conf.Config) (*redis.Client, error) {
	once.Do(func() {
		// 创建Redis客户端对象
		redisClient = redis.NewClient(&redis.Options{
			Addr: config.RedisConfig.Address,
			// 如果有密码，请填写密码
			Password: config.RedisConfig.Password,
			// 选择要使用的数据库编号
			DB: config.RedisConfig.Db,
		})
	})
	// 测试连接是否成功
	err := redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}
