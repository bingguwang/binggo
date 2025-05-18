package model

import (
    "fmt"
    "time"
)

// Consumer 消费者端有缓冲区
type Consumer struct {
    ConsumerChannel chan int
    IsStop          chan bool
}

func NewConsumer(length int) *Consumer {
    return &Consumer{
        ConsumerChannel: make(chan int, length),
        IsStop:          make(chan bool),
    }
}

func (c *Consumer) StopConsume() {
    c.IsStop <- true
}

func (c *Consumer) Consume() {
    go func() {
        for {
            if len(c.ConsumerChannel) == ConsumerBufferLength {
                fmt.Println("-------------消费者队列满了---------------------")
            }
            select {
            case v := <-c.ConsumerChannel:
                time.Sleep(100 * time.Millisecond)
                fmt.Println("消费掉:   ", v)

            case <-c.IsStop:
                fmt.Println("消费停止")
                return
            }
        }
    }()
}

//
//func (c *Consumer) SyncJob(producerChannel chan int) {
//    go func() {
//        for {
//            select {
//            case v := <-producerChannel:
//                c.ConsumerChannel <- v
//            case <-c.IsStop:
//                fmt.Println("消费者停止消费，停止同步操作")
//                return
//            }
//        }
//    }()
//
//}
