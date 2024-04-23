package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// 消费者要怎么处理消息就在这里
	fmt.Printf("从%s消费消息；%s\n", m.NSQDAddress, time.Now().String())

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return nil
}

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic", "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	// 为消费者收到的消息设置handler
	consumer.AddHandler(&myMessageHandler{})
	//consumer.AddConcurrentHandlers 可以同时开启多个协程执行handler处理消息

	// 通过 nsqlookupd 来发现nsq实例
	err = consumer.ConnectToNSQLookupd("192.168.2.130:41610") // 传入的地址是nsqlookupd的地址
	if err != nil {
		log.Fatal(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}
