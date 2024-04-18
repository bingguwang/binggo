package channelP_C

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"testing"
	"time"
)

/**
一写一读
*/
// 一个生产者，一个消费者
func TestCase1(t *testing.T) {
	ch := make(chan int, 5)

	go func() { // 消费者goroutine
		v := <-ch
		fmt.Println("消费了:", v)
	}()
	go func() { // 生产者goroutine
		ch <- 1   // 生产数据并发送到通道
		close(ch) // 关闭通道，表示生产结束
	}()

	osch := make(chan os.Signal, 1)
	signal.Notify(osch, syscall.SIGTERM, syscall.SIGINT)
	<-osch
}

/**
一写多读
*/
// 一个生产者，多个消费者
// 如果要close channel，直接close 生产者，然后通知到消费者
func TestCase2(t *testing.T) {
	ch := make(chan int, 5)
	for i := 0; i < 10; i++ {
		go func() { // 消费者goroutine
			if v, ok := <-ch; ok {
				fmt.Println("消费了:", v)
			} else {
				fmt.Println("生产者退出， channel 关闭")
			}
		}()
	}

	go func() { // 生产者goroutine
		for i := 0; i < 5; i++ {
			ch <- i // 生产数据并发送到通道
		}
		close(ch) // 关闭通道，表示生产结束
	}()

	osch := make(chan os.Signal, 1)
	signal.Notify(osch, syscall.SIGTERM, syscall.SIGINT)
	<-osch
}

/**
多写一读
*/
// 多个生产者，一个消费者
// 未做处理，消费者一股脑的消费
func TestCase3(t *testing.T) {
	ch := make(chan int, 5)
	go func() { // 消费者goroutine
		for v := range ch {
			fmt.Println("消费：", v)
		}
		// ch close后会退出循环
	}()

	// 多个生产者
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i
		}(i)
	}
	wg.Wait()
	close(ch)

	osch := make(chan os.Signal, 1)
	signal.Notify(osch, syscall.SIGTERM, syscall.SIGINT)
	<-osch
}

/*
*
多读多写
多个生产者多个消费者
*/
func TestCase4(t *testing.T) {
	/**
	try-send 是写入channel时候，channel满了我们也不阻塞，加上default分支继续操作
	try-receive 是写入channel时候，channel空了我们也不阻塞，加上default分支继续操作
	*/
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int)
	stopCh := make(chan struct{}) // stopCh 是额外引入的一个信号 channel.

	// 它的生产者是下面的 toStop channel，
	// 消费者是上面 dataCh 的生产者和消费者
	toStop := make(chan string, 1) // 停止信号
	// toStop 是拿来关闭 stopCh 用的，由 dataCh 的生产者和消费者写入
	// 由下面的匿名中介函数(moderator)消费
	// 要注意，这个一定要是 buffered channel （否则没法用 try-send 来处理了）

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop // 收到通知信号后关闭stopCh
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 { // 假设每个生产者如果生成了随机数是0 ，生产者就通知生产
					// try-send 操作
					// 如果 toStop 满了，就会走 default 分支啥也不干，也不会阻塞
					select {
					case toStop <- "sender#" + id: // 生产者写入toStop channel
					default:
					}
					return
				}

				// try-receive 操作，尽快退出
				// 如果没有这一步，下面的 select 操作可能造成 panic
				select {
				case <-stopCh:
					return
				default:
				}

				// 如果尝试从 stopCh 取数据的同时，也尝试向 dataCh
				// 写数据，则会命中 select 的伪随机逻辑，可能会写入数据
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// 同上
				select {
				case <-stopCh:
					return
				default:
				}

				// 尝试读数据
				select {
				case <-stopCh: // 读取到停止信号结束协程
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
					fmt.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	wgReceivers.Wait()
	fmt.Println("stopped by", stoppedBy)
}
