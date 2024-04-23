package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	var topic string

	flag.StringVar(&topic, "topic", "test", "topic name default test")
	flag.Parse()

	// 为每个包含topic的nsqd节点 创建1个生产者
	NewTopicProducer(topic)

	go func() {
		timerTicker := time.Tick(2 * time.Second)
		var i = 1
		for {
			<-timerTicker
			totalNode, failedNode, err := PublishTopicMsg(topic, []byte("消息:"+strconv.Itoa(i)+"  "+time.Now().Format("2006-01-02 15:04:05")))
			if err != nil {
				log.Fatalln("PublishTopicMsg err topic", topic, "err", err.Error())
			}
			log.Println("PublishTopicMsg ok topic", topic, "totalNode", totalNode, "failedNode", failedNode)
			i++

		}
	}()

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sigMsg := <-sigChan
	log.Println("sigMsg", sigMsg)

	// Gracefully stop the producer.
	for _, producers := range TopicProducers {
		for _, producer := range producers {
			producer.Stop()
		}
	}
}

// NewTopicProducer
// 获取 topic 所有的 nsqd 节点 并建立 tcp 链接
func NewTopicProducer(topic string) {
	TopicProducers = make(map[string][]*nsq.Producer)
	config := nsq.NewConfig()
	topicNodeAddr := getTopicNodeAddrSet(topic) // 通过lookup获取 topic 的所在的 nsqd 节点
	var producers []*nsq.Producer
	if len(topicNodeAddr) == 0 { // 一开始所有nsqd上都没有此topic，但是topic又必须在投递消息时候创建
		// 所有nsqd都没有此topic就在所有nsqd节点都创建此topic
		topicNodeAddr = getTopicNodeAddrSet("")
	}
	for _, addr := range topicNodeAddr {
		producer, err := nsq.NewProducer(addr, config)
		if err != nil {
			log.Fatalln("newProducer err topic", topic, "err", err.Error())
		}
		producer.SetLogger(&nsqServerLogger{}, nsq.LogLevelDebug)
		producers = append(producers, producer)
	}
	TopicProducers[topic] = producers
}

// PublishTopicMsg
// 向 topic 发送消息 会自动向每一个包含此 topic 的节点发送 即集群模式
func PublishTopicMsg(topic string, msg []byte) (totalNode int, failedNode int, err error) {
	producers, ok := TopicProducers[topic]
	if !ok {
		return 0, 0, errors.New("PublishTopicMsg err topic not exists")
	}
	totalNode = len(producers)
	for _, producer := range producers {
		errPub := producer.Publish(topic, msg)
		if nil != errPub {
			failedNode++
		}
	}
	return
}

// 获取 topic 的所在的 nsqd 节点集合
func getTopicNodeAddrSet(topic string) (topicNodeAddrArr []string) {
	// 访问 nsqlookupd 的HTTP接口
	url := "http://192.168.2.130:41610/lookup?topic=" + topic
	if topic == "" { // 不用topic过滤，返回所有的nsqd节点
		url = "http://192.168.2.130:41610/nodes"
	}
	resp, _ := http.Get(url)
	defer func() {
		_ = resp.Body.Close()
	}()

	bodyRaw, _ := ioutil.ReadAll(resp.Body)
	lookupTopicRes := &LookupTopicRes{}
	_ = json.Unmarshal(bodyRaw, &lookupTopicRes)

	for _, producer := range lookupTopicRes.Producers {
		topicNodeAddrArr = append(topicNodeAddrArr, producer.BroadcastAddress+":"+strconv.Itoa(producer.TcpPort))
	}

	for _, channel := range lookupTopicRes.Channels {
		fmt.Println("channel:", channel)
	}

	fmt.Println("topicNodeAddrArr：", topicNodeAddrArr)

	return topicNodeAddrArr
}
