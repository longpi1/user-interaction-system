package conf

import (
	"os"

	"github.com/longpi1/gopkg/libary/log"
	"github.com/longpi1/gopkg/libary/queue"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var conf *WebConfig

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

type WebConfig struct {
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
	QueueConfig struct {
		TopicName string       `json:"topic_name"`
		Config    queue.Config `json:"config"`
	} `json:"comment" mapstructure:"comment"`
	AppConfig struct {
		Port        string `json:"port"`
		Debug       bool   `json:"debug"`
		LogFilePath string `json:"log_path" mapstructure:"log_path"`
	} `json:"app" mapstructure:"app"`
}

type RocketConf struct {
	Address  []string `json:"address"`
	LogLevel string   `json:"logLevel"`
}

type PulsarConf struct {
	Address  []string `json:"address"`
	LogLevel string   `json:"logLevel"`
}

type KafkaConf struct {
	Address       []string `json:"address"`
	Version       string   `json:"version"`
	RandClient    bool     `json:"randClient"`
	MultiConsumer bool     `json:"multiConsumer"`
}

type MapConfig struct {
	TypeMap     map[string]int `json:"type_map"  mapstructure:"type_map"`
	PlatformMap map[string]int `json:"platform_map"  mapstructure:"platform_map"`
}

func GetConfig() *WebConfig {
	if conf == nil {
		filePath := getFilePath()
		// 初始化web配置文件
		initWebConfig(filePath)
	}
	return conf
}

func getFilePath() string {
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
	return appFullPath
}

func initWebConfig(appFullPath string) {
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
