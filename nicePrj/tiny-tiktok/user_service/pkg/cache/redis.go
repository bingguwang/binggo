package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	log "tiny-tiktok/user_service/pkg/logger"
)

var Redis *redis.Client

var Ctx = context.Background()

// InitRedis 连接redis
func InitRedis() {
	addr := viper.GetString("redis.address")
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0, // 存入DB0
	})
	log.Log.Info("redis初始化完毕!")
	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
