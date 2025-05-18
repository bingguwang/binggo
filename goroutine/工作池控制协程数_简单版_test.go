package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		results <- task * 2 // 假设任务是将数字乘以 2
	}
}
func TestName(t *testing.T) {
	numWorkers := 4 // 工作池大小
	tasks := make(chan int, 100)
	results := make(chan int, 100)

	// 启动固定数量的 Goroutine
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, tasks, results)
		}(i)
	}

	// 提交任务到 channel
	for i := 1; i <= 20; i++ {
		tasks <- i
	}
	close(tasks)

	// 收集结果
	go func() {
		wg.Wait()
		fmt.Println("关闭results")
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}
