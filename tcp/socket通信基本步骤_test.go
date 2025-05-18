package tcp

import (
	"fmt"
	"net"
	"testing"
)

func TestSocket(t *testing.T) {
	// 创建socket
	serverAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	listener, _ := net.ListenTCP("tcp", serverAddr)

	// 绑定socket并监听端口
	fmt.Println("listening on", serverAddr.String())

	// 等待客户端连接
	for {
		conn, _ := listener.Accept() // 没有连接会一直阻塞，直到有连接才会返回conn
		fmt.Println("client connected")

		// 发送数据
		conn.Write([]byte("Hello, client"))

		// 接收数据
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println("received message:", string(buf[:n]))
		// 关闭连接
		conn.Close()
	}

}
