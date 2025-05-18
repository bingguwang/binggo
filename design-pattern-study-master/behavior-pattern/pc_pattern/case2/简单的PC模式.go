package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	workchan := make(chan int, 10) // 工作队列，10长度

	var wg sync.WaitGroup

	//生产者启动
	wg.Add(1)
	go func(workchan chan<- int) { // 生产者负责放入，所以是只写
		defer wg.Done()
		for {
			ticker := time.NewTicker(time.Second)
			workchan <- time.Now().Second()
			select {
			case v := <-ticker.C:
				if len(workchan) == 10 {
					fmt.Println("满了")
				}
				fmt.Println("time:", v, " 生产:", v)
			}
		}
	}(workchan)

	//消费者启动
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			ticker := time.NewTicker(2 * time.Second)
			val := <-workchan
			select {
			case v := <-ticker.C:
				fmt.Println("time:", v, " 消费:", val)
			}
		}
	}()

	wg.Wait()
	close(workchan)
}
