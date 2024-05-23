package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// 用于处理消息的 processor，需要实现 go-nsq 中定义的 msgProcessor interface，核心是实现消息回调处理方法： func HandleMessage(msg *nsq.Message) error
type msgProcessor struct {
	// 消费者名称
	consumerName string
	// 消息回调处理函数
	callback func(consumerName string, msg []byte) error
}

func newMsgProcessor(consumerName string, callback func(consumerName string, msg []byte) error) *msgProcessor {
	return &msgProcessor{
		consumerName: consumerName,
		callback:     callback,
	}
}

func handleMessageCallback(consumerName string, msg []byte) error {
	fmt.Println(consumerName, "消费了", string(msg))
	return nil
}

// 消息回调处理
func (m *msgProcessor) HandleMessage(msg *nsq.Message) error {
	// 执行用户定义的业务处理逻辑
	if err := m.callback(m.consumerName, msg.Body); err != nil {
		return err
	}
	// 倘若业务处理成功，则调用 Finish 方法，发送消息的 ack
	msg.Finish()
	return nil
}

/*
*
注意，消费者是订阅方，订阅的频道就是 topic+channel名称
所有订阅了同样的频道的consumer，都可以从此频道里获取消息，但是每条消息是会随机推给这些订阅了的某个consumer（即分摊消息）
这就是所谓的一个channel对应多个consumer的正确理解
一个consumer如果只想独享一个topic的全部消息数据，那可以建一个channel名称唯一的consumer来消费
*/
func main() {
	config := nsq.NewConfig()
	// consumer创建时指定了要消费的topic，并且要从名称为MyChannel的channel里获取
	// MyChannel里会有一份完整的topic的消息数据拷贝！
	consumer, err := nsq.NewConsumer("topic", "MyChannel", config)
	if err != nil {
		log.Fatal(err)
	}

	// 为消费者收到的消息设置handler
	processor := newMsgProcessor("C1", handleMessageCallback)
	consumer.AddHandler(processor)
	//consumer.AddConcurrentHandlers 可以同时开启多个协程执行handler处理消息

	// 通过 nsqlookupd 来发现nsq实例
	// consumer需要从topic对应的某个channel里获取消息
	// 每个channel都有一份topic消息的完整拷贝
	// 而topic的载体是一个个的nsqd，所以对于每个consumer需要先得到一个nsqd(有对应topic的nsqd)，然后是channel
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
