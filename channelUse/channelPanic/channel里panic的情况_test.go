package channelPanic

import (
	"fmt"
	"testing"
	"time"
)

// 向关闭的channel发数据会导致panic
// 最好在发数据时，考虑一下是否在其他地方关闭了此channel
// 一个好的原则就是：不要在消费者端close channel
func TestCase1(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Panic:", e)
		} else {
			fmt.Println("没有panic")
		}
	}()
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	ch <- 1
}

// 关闭一个已经关闭的channel会导致panic
// 由于channel是否关闭不好感知到，所以除了用close channel来传递通知消息之外，一般不close channel
// 一个好的原则是： 有多个并发写的生产者时，不要close channel
func TestCase2(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Panic:", e)
		} else {
			fmt.Println("没有panic")
		}
	}()
	ch := make(chan int)

	go func() {
		close(ch)
	}()
	time.Sleep(2 * time.Second)
	close(ch)
}
