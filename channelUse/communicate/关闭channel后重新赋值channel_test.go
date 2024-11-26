package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

/*
*

	假设设置一个channel变量，在开启一些协程之后，给他赋了一个新的channel值，
	那此时的channel 已经是一个全新的channel了，与之前的channel无关
*/

type TestChan struct {
	ch chan string
}

func TestNewChanValue(t *testing.T) {
	stopChap := make(chan os.Signal, 1)
	signal.Notify(stopChap, syscall.SIGTERM, syscall.SIGINT)

	ins := TestChan{ch: make(chan string)}
	go func(ch chan string) {
		time.Sleep(time.Second)
		fmt.Printf("%p\n", ins.ch)
		fmt.Printf("%p\n", ch)
		select {
		case <-ch: // 如果换成 ins.ch那就可以获取到关闭通知，因为ins.ch指向的同一个
			fmt.Println("关闭通道")
			return
		}
	}(ins.ch)

	ins.ch = make(chan string) // 重新赋值，之前的channel和这个新的channel是无关 的
	fmt.Printf("%p\n", ins.ch)
	close(ins.ch)

	<-stopChap
}
