package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

/*
由于没有resp.Body.Close()，泄漏是一定的
ioutil.ReadAll()，但如果此时忘了 resp.Body.Close()，确实会导致泄漏。
因为这里的6次调用的域名一直是同一个的话，那么只会泄漏一个 读goroutine 和一个写goroutine，

所以下面输出的 协程数是3，一个读协程，一个写协程，一个main协程
*/
func main() {
	num := 6
	var resp *http.Response
	for index := 0; index < num; index++ {
		resp, _ = http.Get("https://www.baidu.com")
		_, _ = ioutil.ReadAll(resp.Body)
	}
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}
