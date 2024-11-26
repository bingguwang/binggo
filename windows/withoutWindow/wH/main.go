package main

import "net/http"

func main() {
	http.ListenAndServe(":8989", nil)
}

/**
使用 go build -o myapp.exe -ldflags="-H=windowsgui" 利用通道优雅的关闭gin服务_test.go

*/
