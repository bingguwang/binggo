package channelUse

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

/*
通道可以用于超时处理
*/

// 简单的一个例子
func TestName(t *testing.T) {
	ch := make(chan string)

	// 假如一个任务执行需要2s
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // 超时
		fmt.Println("timeout")
	}
}

// 再来个复杂点的案例，比如http的超时处理
func TestTimeoutHttp(t *testing.T) {
	// 假如我有这样的一堆网址
	urls := []string{
		"https://www.baidu.com",
		"https://www.invalid-url.com", // 无效的URL，模拟超时
	}
	// 创建channel，用于接收协程里的错误信息
	errCh := make(chan string, len(urls))

	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		// 开启协程去访问网站, 如果访问超时了，我们需要在这里收到相关的超时通知
		go func(errCh chan string, url string) {
			defer wg.Done()
			now := time.Now()
			// 访问网站
			client := http.Client{
				Timeout: time.Second * 2, // 设置超时时间为2秒
			}
			resp, err := client.Get(url)
			if err != nil {
				// 超时了放入
				errCh <- fmt.Sprintf("Error fetching %s: %v", url, err)
				return
			}
			defer resp.Body.Close()
			duration := time.Since(now)
			errCh <- fmt.Sprintf("访问%s成功,用时:%s", url, duration)
		}(errCh, urls[i])
	}

	// 使用另外一个协程等待所有协程完成
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(errCh) // 等待所有协程完成后关闭通道
	}(&wg)

	// 会在读取到零值之前把所有内容读完
	for msg := range errCh {
		fmt.Println(msg)
	}
}
