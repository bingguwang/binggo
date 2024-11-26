package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
)

type message struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Payload string `json:"message"`
}

func main() {
	pusher := kq.NewPusher([]string{
		"192.168.2.130:9092",
		"192.168.2.130:9093",
		"192.168.2.130:9095",
	}, "courseware-logDemo") // 创建了一个叫kq的topic

	ticker := time.NewTicker(time.Millisecond)
	for round := 0; round < 3; round++ {
		<-ticker.C

		count := rand.Intn(100)
		m := message{
			// 这里不写key没问题，其实在真正调用push的时候，会创建一个这样的key，用来go-zero里封装实现的分区负载均衡
			Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
			Value:   fmt.Sprintf("%d,%d", round, count),
			Payload: fmt.Sprintf("%d,%d", round, count),
		}
		body, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
		// 推送消息
		if err := pusher.Push(string(body)); err != nil {
			log.Fatal(err)
		}
	}

	cmdline.EnterToContinue()
}
