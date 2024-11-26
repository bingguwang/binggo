package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
	"time"
)

func main() {
	// 创建一个文件保存内存分析结果
	/*	f, err := os.Create("./memprofile.prof")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()*/

	// 定期触发垃圾回收并写入内存分析结果
	go func() {
		for {
			//runtime.GC() // todo 如果不注释，就会触发gc，分析图里可以看到内存占用是0哦
			//if err := pprof.WriteHeapProfile(f); err != nil {
			//	log.Fatal("could not write memory profile: ", err)
			//}
			time.Sleep(1 * time.Minute) // 根据需要调整频率
		}
	}()

	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGTERM, syscall.SIGINT)

	// 你的程序逻辑
	var a []string
	for {
		select {
		case <-time.After(100 * time.Millisecond):
			fmt.Println("写入")
			a = append(a, "aa")
		//size(a)
		case <-osChan:
			fmt.Println("关闭")
			// 在程序结束前创建内存配置文件
			f, err := os.Create("memprofile.prof")
			if err != nil {
				log.Fatal("could not create memory profile: ", err)
			}
			defer f.Close()

			runtime.GC() // 显式调用垃圾回收
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("could not write memory profile: ", err)
			}
			return
		}
	}
}

func size(a []string) {
	fmt.Println(len(a))
}
