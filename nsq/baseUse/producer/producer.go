package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

// 向指定节点的topic发消息，此时不会创建channel， channel 是在有消费者连接并订阅了主题时才会被创建。
func main() {
	nsqAddress := "192.168.2.130:41501"
	nsqAddress2 := "192.168.2.130:41500"

	config := nsq.NewConfig()

	producer, err := nsq.NewProducer(nsqAddress, config)
	if err != nil {
		fmt.Println(err)
	}
	defer producer.Stop()
	producer2, err := nsq.NewProducer(nsqAddress2, config)
	if err != nil {
		fmt.Println(err)
	}
	defer producer2.Stop()
	msg := "测试的内容"
	for {

		if err = producer.Publish("topic", []byte(msg)); err != nil {
			fmt.Printf("producer1 publish message failed, err: %v\n", err)
		} else {
			fmt.Println("msg send to topic (producer1)")
		}

		if err = producer2.Publish("topic", []byte(msg)); err != nil {
			fmt.Printf("producer2 publish message failed, err: %v\n", err)
		} else {
			fmt.Println("msg send to topic (producer2)")
		}
		time.Sleep(10 * time.Second)
	}

}
