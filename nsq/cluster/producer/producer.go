package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"strings"
)

var TopicProducers map[string][]*nsq.Producer // 每个topic对应多个producer

type LookupTopicRes struct {
	Channels  []string       `json:"channels"`
	Producers []ProducerInfo `json:"producers"`
}

type ProducerInfo struct {
	RemoteAddress    string `json:"remote_address"`
	Hostname         string `json:"hostname"`
	BroadcastAddress string `json:"broadcast_address"`
	TcpPort          int    `json:"tcp_port"`
	HttpPort         int    `json:"http_port"`
	Version          string `json:"version"`
}

// nsq 内部日志
type nsqServerLogger struct {
}

func (nsl *nsqServerLogger) Output(callDepth int, s string) error {
	log.Println("nsqServerLogger", callDepth, s[:3], strings.Trim(s[3:], " "))
	return nil
}
