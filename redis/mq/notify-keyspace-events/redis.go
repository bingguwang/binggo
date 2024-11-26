package main

import (
	"context"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx := context.Background()
	fmt.Println(rdb.Ping(ctx).Err())

	go func() {
		var i = 0
		for {
			select {
			case <-time.After(1 * time.Second):
				rdb.Set(ctx, "my_key", "xxx"+strconv.Itoa(i), 0)
			}
			i++
		}
	}()
	//订阅channel
	pubSub := rdb.Subscribe(ctx, "__keyspace@0__:my_key")
	defer pubSub.Close()

	for v := range pubSub.Channel() {
		fmt.Println("msg:", v)
	}

}
