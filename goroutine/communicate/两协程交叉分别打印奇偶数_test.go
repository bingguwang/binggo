package communicate

import (
	"fmt"
	"sync"
	"testing"
)

/*
两个goroutine 交替有序打印 0-100.

这个考察对协程之间通信方式的熟练度
*/

// 方法一
// mutex+cond实现
func TestName(t *testing.T) {
	var (
		wg   sync.WaitGroup
		lock sync.Mutex
		cond = sync.NewCond(&lock)
		isOu = true
	)
	wg.Add(2)

	go func() { // 打印偶数
		defer wg.Done()
		for i := 0; i < 100; i++ {
			cond.L.Lock()
			if !isOu { // 不是偶数就阻塞
				cond.Wait()
			}
			cond.L.Unlock()

			if i%2 == 0 { // 偶数就打印
				cond.L.Lock()
				isOu = true
				cond.L.Unlock()
				fmt.Println("g1:", i)
				cond.Signal()
			}
		}
	}()

	go func() { // 打印奇数
		defer wg.Done()
		for i := 0; i < 100; i++ {
			cond.L.Lock()
			if isOu { // 是偶数就阻塞
				cond.Wait()
			}
			cond.L.Unlock()

			if i%2 == 1 { // 奇数就打印
				cond.L.Lock()
				isOu = false
				cond.L.Unlock()
				fmt.Println("g2:", i)
				cond.Signal()
			}

		}
	}()
	wg.Wait()
}

// 方法二
// channel实现通信
func TestName2(t *testing.T) {
	var (
		wg   sync.WaitGroup
		chOu = make(chan int, 1)
		chJi = make(chan int, 1)
	)
	wg.Add(2)
	chOu <- 1
	go func() { // 打印偶数
		defer wg.Done()
		for i := 0; i < 100; i += 2 {
			<-chOu
			fmt.Println("g1:", i)
			chJi <- 1
		}
	}()

	go func() { // 打印奇数
		defer wg.Done()
		for i := 1; i < 100; i += 2 {
			<-chJi
			fmt.Println("g2:", i)
			chOu <- 1
		}
	}()

	wg.Wait()
}
