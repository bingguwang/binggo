package main

import (
	"time"
)

func main() {
	for i := 0; i < 50000; i++ {
		go func() {
			t := time.NewTimer(5 * time.Second)

			for {
				t.Reset(5 * time.Second)
				t.Reset(10 * time.Second)
				time.Sleep(20 * time.Millisecond)
			}
		}()
	}

	<-make(chan struct{})
}
