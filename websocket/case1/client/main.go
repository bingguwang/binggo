package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func main() {
	// 服务器地址，假设服务器运行在本地8080端口
	//serverURL := "ws://localhost:8080/subscribe"

	// 解析服务器URL
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/subscribe"}
	log.Printf("Connecting to %s", u.String())

	// 创建WebSocket连接
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("Dial error: %v", err)
	}
	defer conn.Close()

	// 读取服务器发送的消息
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			// 读取消息
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				return
			}
			// 打印接收到的消息（这里假设消息是JSON格式的库存信息）
			fmt.Printf("Received message: %s\n", message)
		}
	}()

	// 保持连接一段时间（这里设置为无限循环，直到手动中断）
	// 在实际应用中，你可能会有更复杂的逻辑来处理连接保持和断开
	select {}
}
