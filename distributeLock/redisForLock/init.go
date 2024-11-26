package redisForLock

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.0.66:36379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		panic(err.Error())
	}
	fmt.Println("初始化redis client 成功")
}
func GetClient() *redis.Client {
	return redisClient
}
