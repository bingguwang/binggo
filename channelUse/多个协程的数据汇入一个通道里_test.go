package channelUse

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {

}

var n int = 5000000000

// 多协程计算
// 比如并发求和

// 一般求和
func TestAdd(t *testing.T) {
	now := time.Now()
	var res float64 = 0
	for i := 0; i < n; i++ {
		res += float64(i)
	}
	duration := time.Since(now)
	fmt.Println("used:", duration)
	fmt.Println(res)
}

// 并发求和

func TestAdd2(t *testing.T) {
	now := time.Now()
	glen := 4
	var res float64 = 0
	var wg sync.WaitGroup
	ch := make(chan float64, glen)
	seg := n / glen
	for i := 0; i < glen; i++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			fmt.Println("start:", start, "end:", end)
			var r float64
			for i := start; i < end; i++ {
				r += float64(i)
			}
			ch <- r
		}(i*seg, (i+1)*seg)
	}

	go func() {
		wg.Wait()
		close(ch) // 等待所有协程完成后关闭通道
	}()

	for result := range ch {
		res += result
	}

	duration := time.Since(now)
	fmt.Println("used:", duration)
	fmt.Println(res)
}
