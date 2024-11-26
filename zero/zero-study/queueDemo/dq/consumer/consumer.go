package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var ip = "192.168.2.130"

func main() {
	consumer := dq.NewConsumer(dq.DqConf{
		Beanstalks: []dq.Beanstalk{
			{
				Endpoint: ip + ":11300",
				Tube:     "tube1",
			},
			{
				Endpoint: ip + ":11301",
				Tube:     "tube2",
			},
		},
		Redis: redis.RedisConf{
			Host: ip + ":6379",
			Type: redis.NodeType,
		},
	})
	var (
		producerOnce sync.Once
		producer     dq.Producer
	)
	consumer.Consume(func(body []byte) {
		fmt.Println("一旦拿到消息就消费:", string(body))
		i, _ := strconv.ParseInt(string(body), 10, 64)
		if i%2 == 0 { // 对于消息内容是个偶数的消息，需要重新入队里，然后重新在这里消费掉
			fmt.Println("对于消息内容是个偶数的消息，需要重新入队里，然后重新在这里消费掉")
			i++
			producerOnce.Do(func() {
				producer = dq.NewProducer([]dq.Beanstalk{
					{
						Endpoint: ip + ":11300",
						Tube:     "tube1",
					},
					{
						Endpoint: ip + ":11301",
						Tube:     "tube2",
					},
				})
			})
			if _, err := producer.Delay([]byte(strconv.Itoa(int(i))), 5*time.Second); err != nil {
				return
			}
		}
	})

}
