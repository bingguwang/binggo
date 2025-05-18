package main

import (
	"fmt"
	"sync"
)

// Task represents a task to be executed
type Task func()

// GoroutinePool represents a pool of goroutines to execute tasks
type GoroutinePool struct {
	maxGoroutines int       // 协程池容量
	tasks         chan Task // 任务队列
	wg            sync.WaitGroup
	active        int // 当前运行的协程数
	lock          sync.Mutex
	cond          *sync.Cond // 创建一个和该互斥锁关联的cond

}

// NewGoroutinePool creates a new GoroutinePool
func NewGoroutinePool(maxGoroutines int) *GoroutinePool {
	res := &GoroutinePool{
		maxGoroutines: maxGoroutines,
		tasks:         make(chan Task),
	}
	res.cond = sync.NewCond(&res.lock) // 创建一个和该互斥锁关联的cond
	return res
}

// worker is the goroutine that processes tasks
func (p *GoroutinePool) worker() {
	defer p.wg.Done()
	for task := range p.tasks {
		task() // 拿到任务就执行
		p.lock.Lock()
		p.active-- // 任务执行完当前指向协程数减1
		p.cond.Signal()
		p.lock.Unlock()
	}
}

// Submit 提交任务到池
func (p *GoroutinePool) Submit(task Task) {
	p.lock.Lock()
	if p.active < p.maxGoroutines {
		p.wg.Add(1)
		go p.worker() // 开个worker协程去执行任务
		p.active++
	} else {
		// 没有空闲协程时，应该阻塞等到有协程时候继续执行
		p.cond.Wait()
	}
	p.lock.Unlock()

	p.tasks <- task
}

// Shutdown waits for all goroutines to complete and then closes the pool
func (p *GoroutinePool) Shutdown() {
	close(p.tasks)
	p.wg.Wait()
}

// main function to demonstrate the usage of GoroutinePool
func main() {
	pool := NewGoroutinePool(5)

	// Submit some tasks to the pool
	for i := 0; i < 20; i++ {
		i := i // capture the loop variable
		pool.Submit(func() {
			fmt.Printf("Task %d is running\n", i)
		})
	}

	// Shutdown the pool and wait for all tasks to complete
	pool.Shutdown()
	fmt.Println("All tasks completed")
}
