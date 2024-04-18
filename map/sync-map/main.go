package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	count := 0
	a := sync.Map{}
	for i := 0; i < 10000; i++ {
		a.Store(strconv.Itoa(i), i)
	}

	a.Range(func(key, value any) bool {
		count++
		return true
	})
	fmt.Println(count)
}
