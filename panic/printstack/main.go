package main

import (
	"fmt"
	"runtime"
)

func main() {

}

// 可以打印出panic栈信息，定位panic
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println(string(buf[:n]))
}
