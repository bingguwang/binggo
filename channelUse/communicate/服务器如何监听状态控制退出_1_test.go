package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
*
下面是一种常见的主协程和子协程之间通信的方式

main 通过 传递一个可取消的context给子协程，在子协程里发现context被取消了子协程就return
控制main协程的是exitChan，传给子协程，子协程里就可以控制main协程的退出了

#############################

	总结一点就是，让对方能控制自己的退出，就是把自己的把柄给对方，而这个把柄就是context或者channel

在这里，main交给子协程的把柄就是exitChan，子协程可以通过这个把柄控制父协程的退出
子协程自己通过与main协程传递进来的ctx绑定，控制退出

#############################
*/
func TestCase1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exitChan := make(chan string)
	defer close(exitChan)
	fmt.Println(ctx.Value(""))
	var wg sync.WaitGroup
	go func(ctx context.Context, exitChan chan string) { // 模拟服务
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second):
				intn := rand.Intn(10)
				if intn%6 == 0 {
					// 服务退出了通知主线程退出
					exitChan <- fmt.Sprintf("退出")
					return
				} else {
					fmt.Println("online......")
				}
			}
		}
	}(ctx, exitChan)
LOOP:
	for {
		select {
		case reason := <-exitChan: // 监听退出信号通道
			fmt.Println("exit:", reason)
			cancel()
			break LOOP
		}
	}
	wg.Wait()
}
