package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"testing"
	"time"
)

/*
*
一个无缓冲channel，多个地方读取，当有消息传入的时候只会有一个读者可以读取到，其他协程还是阻塞
*/
func TestOnceForMul(t *testing.T) {
	stopChap := make(chan os.Signal, 1)
	signal.Notify(stopChap, syscall.SIGTERM, syscall.SIGINT)

	ch := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case v := <-ch:
					fmt.Println(i, " 获取到msg", v)
					return
				}
			}
		}(i)
	}
	fmt.Println("当前协程数:", runtime.NumGoroutine())
	time.Sleep(100 * time.Millisecond)
	ch <- "hahaha"
	time.Sleep(100 * time.Millisecond)
	fmt.Println("当前协程数:", runtime.NumGoroutine())

	wg.Wait()

	<-stopChap
}
