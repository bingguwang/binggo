package channelUse

import (
	"fmt"
	"testing"
	"time"
)

func TestNamexx(t *testing.T) {

	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
		close(ch) // 关闭 channel
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
