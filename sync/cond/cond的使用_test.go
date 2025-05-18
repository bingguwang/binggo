package cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
*
sync.Mutex和 sync.Cond的区别
sync.Cond 是条件变量, 需要和mutex配合使用
它允许一个 Goroutine 等待某个条件成立，而其他 Goroutine 可以通知等待者条件已经满足
*/
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

	// 等待 ready 状态的 goroutine， 在被唤醒之前会等待在 cond.Wait()
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
