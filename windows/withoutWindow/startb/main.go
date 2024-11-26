package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8989", nil)
}

/**

go build -o myapp.exe 利用通道优雅的关闭gin服务_test.go

*/
