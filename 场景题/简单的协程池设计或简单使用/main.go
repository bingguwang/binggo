package main

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

func main() {

}

// 使用 channel+worker的方式
type Pool struct {
	workChan chan func() // 工作队列
	wg       sync.WaitGroup
}

func newPool(size int) *Pool {
	pool := &Pool{
		workChan: make(chan func(), size),
	}
	// 启动一定数目的worker goroutine
	pool.wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer pool.wg.Done()
			for task := range pool.workChan {
				task() // 执行任务
			}
		}()
	}
	return pool
}

// 提交任务
func (p *Pool) Submit(task func()) {
	p.workChan <- task
}

// 具备超时控制的提交任务
func (p *Pool) SubmitWithTimeout(ctx context.Context, task func()) error {
	select {
	case p.workChan <- task:
		return nil
	case <-ctx.Done():
		return ctx.Err() // 返回超时错误
	}
}

// 关闭协程池（优雅关闭）
func (p *Pool) Close() {
	close(p.workChan)
	p.wg.Wait()
}
func useChannelWorker() {
	pool := newPool(100)
	pool.Submit(func() {
		fmt.Println("任务")
	})
	time.Sleep(time.Second)
	pool.Close()
	select {}
}

// 使用ants工具库作为协程池 【ants是动态池，会根据任务压力自动控制协程的数目】
func useAnts() {
	// 1.创建协程池
	// 2.提交任务

	pool, _ := ants.NewPool(1000, ants.WithPreAlloc(true)) // 创建容量为1000的协程池
	defer pool.Release()                                   // 释放资源

	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		// 提交任务
		_ = pool.Submit(func() { // 任务处理逻辑
			defer wg.Done()
			time.Sleep(1 * time.Second)
		})
	}
	wg.Wait()
}
