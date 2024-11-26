package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestPa(t *testing.T) {
	fmt.Println("Starting main")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
			PrintStackTrace()
		}
	}()
	level1()
	fmt.Println("Ending main")
}

func level1() {
	fmt.Println("Entering level1")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in level1:", r)
			PrintStackTrace()
		}
	}()
	level2()
	fmt.Println("Exiting level1")
}

func level2() {
	fmt.Println("Entering level2")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in level2:", r)
			PrintStackTrace()
		}
	}()
	level3()
	fmt.Println("Exiting level2")
}

func level3() {
	fmt.Println("Entering level3")
	panic("Panic in level3")
	fmt.Println("Exiting level3")
}

func PrintStackTrace() {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, 2*len(buf))
	}
	fmt.Println("Stack trace:\n", string(buf))
}

var utc8 = time.FixedZone("CST-8", 8*3600)

func TestDE(t *testing.T) {
	mp := make(map[string]interface{}, 10)
	mp["sss"] = "lll"
	fmt.Println(mp["Sss"] == nil)
	fmt.Println(mp["Sss"])
}
