package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
	实现要知道singleFlight是一个大概什么样的东西：
将多个相同的请求合并成一个，以减少重复请求对下游服务的压力
而且一个请求得到的结果可以共享

思路
sync包来避免多个协程对同一个资源的重复请求。
在这个实现中，我们将使用一个共享的映射来跟踪当前的请求，并使用sync.Mutex和sync.WaitGroup来管理并发请求。

需要思考如下问题---
每个请求的状态怎么管理
确保对于同一个键的多个请求能够共享结果，而不是重复调用实际的处理函数。
请求结果如何共享，怎么保存
使用call结构体来包含请求的结果和值，通过sync.WaitGroup来等待请求完成。


*/

// call represents a request that is currently in-flight or has completed.
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group manages different keys' calls.
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // 懒加载，此map用于保存请求的状态
}

// Do ensures that the function fn is only called once for a given key at a time.
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		// A request is already in-flight for this key, wait for it to complete.
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	// This is the first request for this key, create a new call and unlock the mutex.
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	// Call the function.
	c.val, c.err = fn()
	c.wg.Done()

	// Remove the call from the map once it has completed.
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}

// A sample function to demonstrate the usage of the Group.
func slowFunction() (interface{}, error) {
	time.Sleep(2 * time.Second)
	fmt.Println("执行了slowFunction")
	return "result", nil
}

func TestSingleFlight(t *testing.T) {
	var g Group

	// Simulate multiple requests for the same key.
	for i := 0; i < 5; i++ {
		go func(i int) {
			val, err := g.Do("key", slowFunction)
			if err != nil {
				fmt.Printf("goroutine %d received an error: %v\n", i, err)
			} else {
				fmt.Printf("goroutine %d received value: %v\n", i, val)
			}
		}(i)
	}

	// Wait to see all outputs.
	time.Sleep(5 * time.Second)
}
