package _select

import (
	"fmt"
	"testing"
	"time"
)

func TestASD(t *testing.T) {
	for {
		select {
		case <-time.After(time.Second * 3):
			fmt.Println("十秒后退出")
			break
		}
		fmt.Println("跳出到这里") // break只会跳出select语句块，不能跳出for循环
	}
}

func TestASD1(t *testing.T) {
out:
	for {
		select {
		case <-time.After(time.Second * 3):
			fmt.Println("十秒后退出")
			break out
		}
		fmt.Println("跳出到这里")
	}
} // break跳出for循环，跳到out，且不会再进入for循环
