package channelMemleak

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

// 无非就是channel没有正确关闭导致 goroutine没有退出

/*
*
生产者退出，导致消费者阻塞
*/
func TestCase2(t *testing.T) {
	ch := make(chan int)

	go func() {
		for v := range ch {
			fmt.Println("消费:", v)
		}
	}()
	printCurrentGoroutineNum()

	ch <- 1
	ch <- 2
	//生产者停止生产，close ch可避免消费者协程泄漏
	printCurrentGoroutineNum()

	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGTERM, syscall.SIGINT)
	<-osChan
}
