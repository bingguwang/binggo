package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestHeap(t *testing.T) {
	printAlloc()
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
