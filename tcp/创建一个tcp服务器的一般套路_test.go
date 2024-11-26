package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
)

/*
创建tcp服务器一般套路就是一个for循环去循环接收和处理请求---也就是`for + listener.accept `模式
tcp服务器需要一个listener
*/
func TestTcpSever(t *testing.T) {

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	for {
		conn, err := listen.Accept() // 一个conn代表一个连接, 每个连接也就是对应了一个四元组
		if err != nil {
			fmt.Println("err:", err.Error())
			continue
		}

		// 处理此连接上收到的客户请求
		go func(conn net.Conn) {
			defer conn.Close()
			fmt.Println("处理请求:", conn)

			// 针对当前连接做发送和接受操作
			// 这里的循环是为了一直接收请
			for {
				reader := bufio.NewReader(conn)
				var buf [128]byte
				n, err := reader.Read(buf[:])
				if err != nil {
					fmt.Printf("read from conn failed, err:%v\n", err)
					break
				}

				recv := string(buf[:n])
				fmt.Printf("收到的数据：%v\n", recv) // 注意，每次读出来的不一定是一个完整的请求

				// 因为这里是简单实现，所以并没有在这里拼接完整的请求
				// 其实应该要判断收到一个完整的请求后才向服务端响应

				// 将接受到的数据返回给客户端
				_, err = conn.Write([]byte("ok"))
				if err != nil {
					fmt.Printf("write from conn failed, err:%v\n", err)
					break
				}
			}
		}(conn)

	}

}

func TestTcpClient(t *testing.T) {
	// 1、与服务端建立连接
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	// 2、使用 conn 连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed, err:%v\n", err)
			return
		}
		// 从服务端接收回复消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed:%v\n", err)
			return
		}
		fmt.Printf("收到服务端回复:%v\n", string(buf[:n]))
	}
}
