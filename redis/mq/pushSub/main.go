package main

import (
	"context"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx := context.Background()
	fmt.Println(rdb.Ping(ctx).Err())

	osStopChan := make(chan os.Signal, 1)
	signal.Notify(osStopChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		//订阅channel
		pubSub := rdb.Subscribe(ctx, "vcsIPC")
		defer pubSub.Close()

		for {
			select {
			case v := <-pubSub.Channel():
				fmt.Println("msg:", v)
			case <-osStopChan:
				return
			}
		}

	}()
	<-osStopChan
}

// message格式: key:key值:|del
// message格式: key:key值:|set
// message格式: key:key值:|add
// message格式: key_hash:key值:filed值|del
// message格式: key_hash:key值:filed值|set
// message格式: key_hash:key值:filed值|add
