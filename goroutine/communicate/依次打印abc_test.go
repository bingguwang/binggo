package communicate

import (
	"fmt"
	"sync"
	"testing"
)

func printA(chA, chB chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-chA // 等待 chA 的信号
		fmt.Print("a")
		chB <- struct{}{} // 发送信号给 chB
	}
}

func printB(chB, chC chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-chB // 等待 chB 的信号
		fmt.Print("b")
		chC <- struct{}{} // 发送信号给 chC
	}
}

func printC(chC, chA chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-chC // 等待 chC 的信号
		fmt.Print("c")
		//if i < 9 { // 无缓冲的channel就得这么处理
		//	chA <- struct{}{} // 发送信号给 chA，继续下一轮
		//}
		chA <- struct{}{} // 发送信号给 chA，继续下一轮

	}
}

func TestNamemm(t *testing.T) {

	chA := make(chan struct{}, 1)
	chB := make(chan struct{}, 1)
	chC := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(3)

	go printA(chA, chB, &wg)
	go printB(chB, chC, &wg)
	go printC(chC, chA, &wg)

	// 启动流程
	chA <- struct{}{}

	wg.Wait()
	fmt.Println() // 换行
}
