package cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCond1(t *testing.T) {
	var mu sync.Mutex
	cond := sync.NewCond(&mu) // 创建一个和该互斥锁关联的cond
	ready := false

	// 准备工作，改变 ready 状态并唤醒所有等待的 goroutine
	go func() {
		fmt.Println("准备工作中")
		time.Sleep(2 * time.Second) // 模拟准备工作需要的时间
		mu.Lock()
		ready = true
		fmt.Println("准备好了!")
		cond.Broadcast() // 唤醒所有等待的 goroutine
		mu.Unlock()
	}()

	// 等待 ready 状态的 goroutine
	wait := func(id int) {
		mu.Lock()
		for !ready {
			fmt.Println(id, "等待，等待准备工作结束")
			cond.Wait() // 等待条件满足
		}
		fmt.Printf("Goroutine %d is running\n", id)
		mu.Unlock()
	}

	// 启动多个 goroutine，等待条件满足
	for i := 0; i < 3; i++ {
		go wait(i)
	}

	// 等待一段时间，以便所有 goroutine 执行完成
	time.Sleep(5 * time.Second)
}
