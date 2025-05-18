package model

import "fmt"

type Syncer struct {
    Produce  *Producer
    Consumer *Consumer
}

func NewSyncer(producer *Producer, consumer *Consumer) *Syncer {
    return &Syncer{
        Produce:  producer,
        Consumer: consumer,
    }
}

func (c *Syncer) SyncJob() {
    go func() {
        for {
            select {
            case v := <-c.Produce.ProduceChannel:
                c.Consumer.ConsumerChannel <- v
            case <-c.Consumer.IsStop:
                fmt.Println("消费者停止消费，停止同步操作")
                return
            }
        }
    }()

}
