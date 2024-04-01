package conf

import (
	"os"
	"relation-service/libary/log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var conf *Config

const (
	SchemaTypeDev    = "dev"
	SchemaTypeOnline = "online"
)

const (
	SchemaPathDev    = "dev/"
	SchemaPathOnline = "online/"
)

// 数据库类型
const (
	TypeMySql      = "mysql"
	TypePostgreSQL = "pgsql"
	TypeMSSQL      = "mssql"
)

type Config struct {
	DBConfig struct {
		Type string `json:"type"`
		// 最高优先级
		Dsn string `json:"dsn"`
	} `json:"db"  mapstructure:"db"`
	RedisConfig struct {
		Address  string `json:"addr"`
		Db       int    `json:"db"`
		Password string `json:"password"`
	} `json:"redis" mapstructure:"redis"`
	LocalCache struct {
		EvictionTime time.Duration `json:"eviction_time" mapstructure:"eviction_time"`
	} `json:"local_cache" mapstructure:"local_cache"`
	AppConfig struct {
		Port        string `json:"port"`
		Debug       bool   `json:"debug"`
		LogFilePath string `json:"log_path" mapstructure:"log_path"`
	} `json:"app" mapstructure:"app"`
}

func GetConfig() *Config {
	return conf
}

func init() {
	var filePath string
	// 通过环境变量获取开发模式
	schema := os.Getenv("schema")

	switch schema {
	case SchemaTypeDev:
		filePath = SchemaPathDev
	case SchemaTypeOnline:
		filePath = SchemaPathOnline
	default:
		filePath = SchemaPathDev
	}

	appFullPath := "./conf/" + filePath
	// 解析 config
	viper.SetConfigName("web")
	viper.AddConfigPath(appFullPath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("解析文件失败: ", err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal("解析文件失败: ", err)
	}
	// 监听配置更新
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&conf); err != nil {
			log.Fatal("解析文件失败: ", err)
		}
	})
}
