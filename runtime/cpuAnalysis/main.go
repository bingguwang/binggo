package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"
)

/*
*
本案例中介绍了如何使用 pprof来进行cpu分析的
*/
func main() {
	// 创建porf文件
	f, err := os.Create(`C:\\Users\\dell\\3D Objects\\binggo\\runtime\\pprof\\a.prof`)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGTERM, syscall.SIGINT)

	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println(err)
		return
	}

	//t := time.NewTicker(time.Millisecond * 50)
loop:
	for {
		select {
		case <-osChan:
			pprof.StopCPUProfile()
			break loop
		case v := <-time.NewTicker(time.Second * 5).C:
			fmt.Println(v.String())
			time.Sleep(300 * time.Millisecond)
			//case <-t.C:
		}
	}

	//  启动后最好运行时间就一些，12小时以上
	//  然后关闭程序，此时prof文件里就会有内容了
	//	go tool pprof -http=:8080 a.prof ，然后去到此网址看分析结果
}
