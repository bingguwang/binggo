package model

import (
    "fmt"
    "time"
)

type Producer struct {
    ProduceChannel chan int
}

func NewProducer(length int) *Producer {
    return &Producer{
        ProduceChannel: make(chan int, length),
    }
}

func (p *Producer) Produce(productLength int) {
    go func() {
        // 生产
        for i := 0; i < 200; i++ {
            if len(p.ProduceChannel) == ProducerBufferLength {
                fmt.Println("-------------生产者队列满了---------------------")
            }
            time.Sleep(50 * time.Millisecond)
            select {
            case p.ProduceChannel <- i:
                fmt.Println("生产:", i)
                // 生产者队列满了会阻塞在这
            }
        }
    }()
}
