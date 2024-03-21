package channelMemleak

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"testing"
	"time"
)

// 无非就是channel没有正确关闭导致 goroutine没有退出

/*
*
消费者退出，导致生产者阻塞
*/
func TestCase1(t *testing.T) {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second) // 消费者退出之后，生产者还在sleep
		ch <- 1                 // 消费者退出，这里阻塞
	}()
	printCurrentGoroutineNum()

	select {
	case v := <-ch:
		fmt.Println("消费:", v)
	case <-time.After(time.Millisecond * 500):
		fmt.Println("消费者退出")
	}

	printCurrentGoroutineNum() // 和上面输出一样，生产者协程未退出

	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGTERM, syscall.SIGINT)
	<-osChan
}

func printCurrentGoroutineNum() {
	fmt.Println("当前协程数:", runtime.NumGoroutine())
}
