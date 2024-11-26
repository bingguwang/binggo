package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"testing"
)

/*
*
下面的通信有这几处：
 1. 主线程通过cancel 通知 子http服务的 ctx.Done()
 2. 子http服务通过exitChan通道通知主线程的 <-exitChan
*/
func TestCase2(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	exitChan := make(chan string)
	defer close(exitChan)

	wg.Add(1)
	go func(ctx context.Context, exitChan chan<- string, ip string) { // 模拟http服务
		defer wg.Done()

		httpSrv := &http.Server{
			Addr:    ":7777",
			Handler: gin.Default(),
		}

		go func() {
			if err := httpSrv.ListenAndServe(); err != nil {
				// 服务退出，通知退出通道告知主线程
				exitChan <- "Http ListenAndServe:" + err.Error()
			}
		}()
		defer func() {
			if err := httpSrv.Shutdown(ctx); err != nil {
				fmt.Println("Http Shutdown:" + err.Error())
			}
		}()

		fmt.Println("http service start")

		// 这里是确保主线程收到退出通知后，这里就会放行
		<-ctx.Done()

		fmt.Println("http service end")
	}(ctx, exitChan, "127.0.0.1")

LOOP:
	for {
		select {
		case reason := <-exitChan:
			fmt.Println("exit:", reason)
			cancel()
			break LOOP
		}
	}

	wg.Wait()

	fmt.Println("ngw daemon exit")
}
