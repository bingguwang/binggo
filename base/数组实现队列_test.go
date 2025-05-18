package base

import (
	"fmt"
	"sync"
	"testing"
)

type Queue struct {
	data []int // 存储队列数据的切片
	mu   sync.Mutex
}

// 入队操作
func (q *Queue) Enqueue(value int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, value)
}

// 出队操作
func (q *Queue) Dequeue() (int, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) == 0 {
		return 0, false // 队列为空
	}
	value := q.data[0]  // 获取第一个元素
	q.data = q.data[1:] // 移除第一个元素
	return value, true  // 返回值和是否成功标志
}

// 获取队列长度
func (q *Queue) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.data)
}
func TestNsssame(t *testing.T) {
	queue := &Queue{}

	// 生产者：向队列中添加元素
	go func() {
		for i := 1; i <= 5; i++ {
			queue.Enqueue(i)
			fmt.Printf("Producing: %d\n", i)
		}
	}()

	// 消费者：从队列中读取元素
	for {
		if queue.Length() == 0 {
			continue // 如果队列为空，继续等待
		}
		value, ok := queue.Dequeue()
		if ok {
			fmt.Printf("Consuming: %d\n", value)
		}
		if queue.Length() == 0 && value == 5 { // 假设生产者只生产5个元素
			break
		}
	}
}
