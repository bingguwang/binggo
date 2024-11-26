package config

import "github.com/redis/go-redis/v9"

var (
	GVA_REDIS  *redis.Client
	GVA_CONFIG ServerConfig

	//GVA_DB     *gorm.DB
	//GVA_DBList map[string]*gorm.DB
	//// GVA_LOG    *oplogging.Logger
	//GVA_LOG                 *zap.Logger
	//GVA_Timer               timer.Timer = timer.NewTimerTask()
	//GVA_Concurrency_Control             = &singleflight.Group{}
	//
	//BlackCache local_cache.Cache
	//lock       sync.RWMutex
)

type ServerConfig struct {
	Redis *Redis `json:"redis" yaml:"redis"`
}

type Redis struct {
	DB       int    `json:"db" yaml:"db"`             // redis的哪个数据库
	Addr     string `json:"addr" yaml:"addr"`         // 服务器地址:端口
	Password string `json:"password" yaml:"password"` // 密码
}
