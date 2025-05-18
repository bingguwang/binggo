package channelUse

import (
	"fmt"
	"testing"
	"time"
)

func TestNamexxx(t *testing.T) {

	// 创建一个带缓冲的channel，最大容量为3
	queue := make(chan int, 3)

	// 生产者 Goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Producing: %d\n", i)
			time.Sleep(100 * time.Millisecond)
			queue <- i // 如果队列已满，会阻塞
		}
		close(queue) // 关闭channel表示生产结束
	}()

	// 消费者主程序
	for data := range queue {
		fmt.Printf("Consuming: %d\n", data)
	}
}
