package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
*
通过chan 通知父协程退出
*/
type gw struct {
	osStopChan chan os.Signal
	exitChan   chan struct{}
}

func (g *gw) Run() {
	defer close(g.exitChan)
	fmt.Println("gw working")
	time.Sleep(5 * time.Second)
}

func main() {

	//信号量注册
	osStopChan := make(chan os.Signal, 1)
	signal.Notify(osStopChan, syscall.SIGTERM, syscall.SIGINT)

	gw := &gw{
		osStopChan: osStopChan,
		exitChan:   make(chan struct{}),
	}

	go gw.Run()

	for {
		select {
		case _, ok := <-gw.exitChan:
			if !ok {
				fmt.Println("gw closed")
				return
			}
		case <-time.After(3 * time.Second):
			fmt.Println("gw is running")
		}
	}

}
