package channelUse

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestNamexxaa(t *testing.T) {
	ch := make(chan int, 1)
	ch1 := make(chan int, 1)
	go d(ch1, ch)

	time.Sleep(time.Second)
	close(ch)

	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		ch1 <- 1
	}

	osch := make(chan os.Signal, 1)
	signal.Notify(osch, syscall.SIGTERM, syscall.SIGINT)
	<-osch
}

func d(ch1, ch chan int) {
	var i int
	for {
		i++
		time.Sleep(200 * time.Millisecond)
		select {
		case <-ch1:
			fmt.Println("ccccccc")
		case <-ch:
			fmt.Println("零值")
		default:
			fmt.Println("xxxx")
		}
	}
}
