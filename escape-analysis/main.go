// package main
//
// import (
//
//	"fmt"
//	"net/http"
//	"os"
//	"os/signal"
//	"runtime/pprof"
//	"syscall"
//	"time"
//
// )
//
//	func main() {
//		f, err := os.OpenFile("C:\\Users\\dell\\3D Objects\\binggo\\escape-analysis\\a.prof", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
//		if err != nil {
//			fmt.Println("could not create memory profile: ", err)
//			return
//		}
//		defer f.Close()
//		osChan := make(chan os.Signal, 1)
//		signal.Notify(osChan, syscall.SIGTERM, syscall.SIGINT)
//		go func() {
//			// 启动一个 goroutine, 不阻止正常代码运行
//			http.ListenAndServe("localhost:6060", nil) // 使用 pprof 监听端口
//		}()
//		// 你的程序逻辑
//
// loop:
//
//	for {
//		select {
//		case <-osChan:
//			if err := pprof.WriteHeapProfile(f); err != nil {
//				fmt.Println("could not write memory profile: ", err)
//			}
//			break loop
//		case <-time.NewTicker(time.Millisecond * 50).C:
//			fmt.Println("写一下")
//
//		}
//	}
//
// }
package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
		// 启动一个 goroutine, 不阻止正常代码运行
		http.ListenAndServe("localhost:6060", nil) // 使用 pprof 监听端口
	}()
	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGTERM, syscall.SIGINT)
loop:
	for {
		select {
		case <-osChan:
			break loop
		case <-time.NewTicker(time.Second * 1).C:
		}
	}
}
