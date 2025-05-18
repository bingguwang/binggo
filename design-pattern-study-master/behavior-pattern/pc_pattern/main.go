package main

import (
	"design-pattern-study-master/behavior-pattern/pc_pattern/model"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	consumer := model.NewConsumer(model.ConsumerBufferLength)
	producer := model.NewProducer(model.ProducerBufferLength)

	// 消费开启
	consumer.Consume()

	// 同步器
	syncer := model.NewSyncer(producer, consumer)
	// 同步数据
	syncer.SyncJob()

	// 生产
	producer.Produce(200)

	chann := make(chan os.Signal)
	signal.Notify(chann, os.Interrupt, syscall.SIGINT)
	<-chann
}
