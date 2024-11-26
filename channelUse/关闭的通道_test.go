package channelUse

import (
	"fmt"
	"testing"
	"time"
)

/*
我们知道

	向关闭的通道读取会得到零值

但是关闭之后，如果通道里有没取完的元素，那么会取完之后才取到零值
*/
func TestCase1(t *testing.T) {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("当前channel内有元素:", len(ch))
	close(ch)
	// 可以看到即使关闭了通道，也是会读取通道内剩余的值，直到读到零值才会结束
	for v := range ch {
		fmt.Println("读取出:", v)
	}
}

func TestCase2(t *testing.T) {
	ch := make(chan int, 1)
	go func() {
		ch <- 1
	}()

	time.Sleep(time.Second)
	fmt.Println("当前channel内有元素:", len(ch))
	close(ch)
	// 可以看到即使关闭了通道，也是会读取通道内剩余的值，直到读到零值才会结束
	for v := range ch {
		fmt.Println("读取出:", v)
	}
}

func TestCase3(t *testing.T) {
	replyCh := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			replyCh <- i
			time.Sleep(time.Millisecond * 100) // 模拟一些延迟
		}
		close(replyCh)
	}()

	/*for reply := range replyCh {
		fmt.Println(reply)
	}*/
	for {
		select {
		case v := <-replyCh:
			if v == 0 {
				fmt.Println("读出零值，通道已关闭", v)
				return
			}
			fmt.Println("读出:", v)
		}
	}
}
