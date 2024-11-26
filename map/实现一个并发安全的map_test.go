package _map

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 不安全的map。会出现 fatal error: concurrent map writes
func TestNoSafe(t *testing.T) {
	mp := make(map[int]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mp[i] = i
		}(i)
	}
	wg.Wait()
	fmt.Println(mp)
}

// 安全的map
func TestSafe(t *testing.T) {
	mp := make(map[int]int, 0)
	var (
		wg   sync.WaitGroup
		lock sync.Mutex
	)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			mp[i] = i
			lock.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(mp)
}

// 不用锁如何解决map并发安全的问题
func TestWithoutLock(t *testing.T) {
	// 创建一个 map，用于存储数据
	data := make(map[string]int)

	// 创建一个用于传递操作的 channel
	commands := make(chan command)

	// 单独的 goroutine 处理 map 的操作
	go func() {
		// channel是并发安全的，其实就是利用channel的并发安全来实现了map的并发安全(channel并发安全还是有通过mutex，其实就其本质离不开mutex)
		for cmd := range commands {
			switch cmd.action {
			case "put":
				data[cmd.key] = cmd.value
				fmt.Printf("Put key '%s' with value %d\n", cmd.key, cmd.value)
			case "get":
				val := data[cmd.key]
				cmd.result <- val
			case "delete":
				delete(data, cmd.key)
				fmt.Printf("Deleted key '%s'\n", cmd.key)
			}
		}
	}()

	// 示例：并发操作 map
	numOperations := 10
	results := make(chan int, numOperations)

	// 并发写入
	for i := 0; i < numOperations; i++ {
		key := fmt.Sprintf("key%d", i)
		commands <- command{action: "put", key: key, value: i}
	}

	// 并发读取
	for i := 0; i < numOperations; i++ {
		key := fmt.Sprintf("key%d", i)
		resultChan := make(chan int)
		commands <- command{action: "get", key: key, result: resultChan}
		go func() {
			val := <-resultChan
			results <- val
		}()
	}

	// 等待所有读取操作完成
	time.Sleep(1 * time.Second)

	// 输出所有读取的结果
	close(results)
	//for result := range results {
	//	fmt.Println("Read value:", result)
	//}
	// for迭代channel其实等价于下面的这种写法
out:
	for {
		select {
		case result, ok := <-results:
			if ok {
				fmt.Println("Read value:", result)
			} else {
				break out
			}
		}
	}

	// 关闭命令通道
	close(commands)
}

// command 结构体，用于传递操作指令和相关参数
type command struct {
	action string   // 操作类型：put、get、delete
	key    string   // 键
	value  int      // 值（用于 put 操作）
	result chan int // 结果通道（用于 get 操作）
}
