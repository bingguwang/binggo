package main

import (
	"fmt"
	"time"
)

func main() {
	c := New(-1, 2*time.Second)
	c.Set("k1", "v1", 2*time.Second)

	for {
		select {
		case <-time.After(time.Second):
			fmt.Println(c.Items())
		}
	}
}
