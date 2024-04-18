package channelMap

import (
	"fmt"
	"testing"
	"time"
)

func TestCase1(t *testing.T) {
	ch := make(chan map[int]int, 10)
	fmt.Println("当前ch内长度", len(ch))

	var toinsert = map[int]int{}
	for i := 0; i < 20; i++ {
		toinsert[i] = i
	}
	go func() {
		for v := range ch {
			time.Sleep(time.Second)
			fmt.Println("consume  ", v)
		}
	}()

	for {
		select {
		case ch <- toinsert:
			fmt.Println("当前ch内长度", len(ch))
		case <-time.After(2 * time.Second):
		}
	}
}
