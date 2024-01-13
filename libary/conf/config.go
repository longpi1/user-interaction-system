package conf


type DBType string

const (
	TypeMySql      DBType = "mysql"
	TypePostgreSQL DBType = "pgsql"
	TypeMSSQL      DBType = "mssql"
)

type DBConf struct {
	Type DBType `koanf:"type" json:"type"`
	Dsn  string `koanf:"dsn" json:"dsn"` // 最高优先级

	File string `koanf:"file" json:"file"`
	Name string `koanf:"name" json:"name"`

	Host     string `koanf:"host" json:"host"`
	Port     int    `koanf:"port" json:"port"`
	User     string `koanf:"user" json:"user"`
	Password string `koanf:"password" json:"password"`

	TablePrefix string `koanf:"table_prefix" json:"table_prefix"`
	Charset     string `koanf:"charset" json:"charset"`
	SSL         bool   `koanf:"ssl" json:"ssl"`
}

type CacheType string

// # Redis 配置
// redis:
//
//	network: "tcp"
//	username: ""
//	password: ""
//	db: 0
type RedisConf struct {
	Network  string `koanf:"network" json:"network"` // tcp or unix
	Username string `koanf:"username" json:"username"`
	Password string `koanf:"password" json:"password"`
	DB       int    `koanf:"db" json:"db"` // Redis 默认数据库 0
}