package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建一个 HTTP 服务器
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}

	// 启动 HTTP 服务
	go func() {
		fmt.Println("Starting HTTP server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Could not start server: %v\n", err)
		}
	}()

	// 创建一个 channel 用于接收信号
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// 等待信号
	<-stop
	fmt.Println("Shutting down server...")

	// 创建一个超时上下文，用于优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server shutdown failed: %v\n", err)
	}

	fmt.Println("Server gracefully stopped")
}

// 简单的 HTTP 处理函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
