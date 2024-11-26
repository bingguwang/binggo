package timer

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"testing"
	"time"

	_ "net/http/pprof"
)

// after是否可以放在循环里
func TestAfter(t *testing.T) {
	go http.ListenAndServe(":8080", nil)
	// 通过pprof分析 cpu以及mem的占用情况, 比较直观的就是看他们cpu的火焰图
	//AnalysisCpu(AfterTest)
	AnalysisCpu(TickTest)
	//AnalysisCpu(TickTestCom)
}

func AfterTest(osChan chan os.Signal) {
	for {
		select {
		case <-time.After(10 * time.Millisecond):
		case <-osChan:
			return
		}
	}
}

func TickTestCom(osChan chan os.Signal) {
	t := time.NewTicker(10 * time.Millisecond) // 正常的用法
	for {
		select {
		case <-t.C:
		case <-osChan:
			return
		}
	}
}

func TickTest(osChan chan os.Signal) {
	for {
		select {
		case <-time.NewTicker(10 * time.Millisecond).C: // 会占用大量cpu资源
		case <-osChan:
			return
		}
	}
}

func AnalysisCpu(function func(osChan chan os.Signal)) {
	// 创建prof文件
	f, err := os.Create(`C:\\Users\\dell\\3D Objects\\binggo\\time-study\\timer\\a.prof`)
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
	function(osChan)

loop:
	for {
		select {
		case <-osChan:
			pprof.StopCPUProfile()
			break loop
		}
	}

	//  关闭程序，此时prof文件里就会有内容了
	//	go tool pprof -http=:8080 a.prof ，然后去到此网址看分析结果
}
