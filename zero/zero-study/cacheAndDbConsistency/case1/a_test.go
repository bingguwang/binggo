package case1

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"testing"
	"time"
)

func queryDatabase(key string) (string, error) {
	fmt.Println("进行数据库查询")
	time.Sleep(2 * time.Second) // 模拟查询数据库的延迟
	return fmt.Sprintf("Select value for %s", key), nil
}

var group singleflight.Group

func TestCase1(t *testing.T) {
	var wg sync.WaitGroup
	key := "exampleKey"
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value, err, _ := group.Do(key, func() (interface{}, error) {
				return queryDatabase(key)
			}) // Do里的func只会有一个协程去做，得到的值所有协程共享
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Got value:", value)
			}
		}()
	}
	wg.Wait()
}
