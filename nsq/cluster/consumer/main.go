package main

import (
	"flag"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type nsqMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *nsqMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// do whatever actual message processing is desired
	log.Println("HandleMessage nsqd:", m.NSQDAddress, "msg:", string(m.Body))

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return nil
}

// 每个topic的channel对应了一个消费者组，组内负载均衡选择一个消费者来消费
// main里创建了一个消费者组来对 此topic+channel对应的消息队列来消费
func main() {
	var topic string
	var channel string
	var count int // 要创建的消费者组里消费者的数量
	var consumerGroup []*nsq.Consumer

	flag.StringVar(&topic, "topic", "test", "topic name default test")
	flag.StringVar(&channel, "channel", "test", "channel name default test")
	flag.IntVar(&count, "count", 1, "consumer count default 1")
	flag.Parse()

	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	config.MaxInFlight = 10 // 一个消费者可同时接收的最多消息数量

	for i := 0; i < count; i++ {
		consumer, err := nsq.NewConsumer(topic, channel, config)
		if err != nil {
			log.Fatalln("NewConsumer err:", err.Error())
		}

		// Set the Handler for messages received by this Consumer. Can be called multiple times.
		// See also AddConcurrentHandlers.
		consumer.AddHandler(&nsqMessageHandler{})
		// 并发模式的消费处理 消费者会启用 n 个协程处理消息
		consumer.AddConcurrentHandlers(&nsqMessageHandler{}, 10)
		consumer.SetLogger(&nsqServerLogger{}, nsq.LogLevelDebug)

		// Use nsqlookupd to discover nsqd instances.
		// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
		// 会订阅所有包含当前 topic 的 nsqd 实例
		// 多用于集群模式时 生产者向多个含有topic的实例同时发送消息
		// 当其中部分实例挂到时 消费者仍可通过其它实例获得消息
		// !此处要做消息幂等处理！
		err = consumer.ConnectToNSQLookupd("192.168.2.130:41610")
		if err != nil {
			log.Fatalln("ConnectToNSQLookupd err:", err.Error())
		} else {
			log.Println("ConnectToNSQLookupd success topic:", topic, "channel:", channel)
		}

		consumerGroup = append(consumerGroup, consumer)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sigMsg := <-sigChan
	log.Println("sigMsg", sigMsg)

	// Gracefully stop the consumer.
	for _, consumer := range consumerGroup {
		consumer.Stop()
	}
}
