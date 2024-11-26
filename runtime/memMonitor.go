package main

import (
	"log"
	"runtime"
	"time"
)

func logMemoryUsage() {
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)
		log.Printf("Alloc = %v MiB", bToMb(m.Alloc))
		log.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
		log.Printf("\tSys = %v MiB", bToMb(m.Sys))
		log.Printf("\tNumGC = %v\n", m.NumGC)
		time.Sleep(1 * time.Minute) // 根据需要调整频率
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	go logMemoryUsage()

	// 你的程序逻辑
}
