package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

/*
常见的负载均衡算法就是轮训

下面模拟nginx的负载均衡分发实现
*/

func startServer(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "由 port:%d 的服务器处理 ", port)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	fmt.Printf("Starting server on port %d\n", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}

// 创建一个轮训调度器
type RoundRobinBalancer struct {
	servers []string // 保存所有的服务器地址
	current int      // 当前选择的服务器地址
}

func NewRoundRobinBalancer(servers []string) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		servers: servers,
		current: 0,
	}
}

// 轮训算法，返回一个被选择的服务器地址
func (r *RoundRobinBalancer) NextServer() string {
	server := r.servers[r.current]
	r.current = (r.current + 1) % len(r.servers)
	return server
}

// 来进行分发, 在nginx的负载均衡里做的就是分发请求的工作，就类似这里
func startRobinLoadBalancer() {

	servers := []string{
		"http://localhost:18081",
		"http://localhost:18082",
		"http://localhost:18083",
	}
	balancer := NewRoundRobinBalancer(servers)

	// 启动负载均衡器
	// 创建负载均衡器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 选择后端服务器
		server := balancer.NextServer()

		// 创建转发请求, 转发给选出来的服务器
		proxyURL, _ := url.Parse(server)
		proxyReq, err := http.NewRequest(r.Method, proxyURL.String()+r.URL.Path, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		proxyReq.Header = r.Header

		// 发起请求
		resp, err := http.DefaultClient.Do(proxyReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// 将响应返回给客户端
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.WriteHeader(resp.StatusCode)
		w.Write(body) // 写回http响应里
	})

	fmt.Println("Starting load balancer on port 8080") // 好比是nginx暴露回给前端的端口一样
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err.Error())
	}

}

func TestServer(t *testing.T) {
	// 启动负载均衡器(nginx)
	go startRobinLoadBalancer()

	// 启动三个模拟的后端服务器
	go startServer(18081)
	go startServer(18082)
	go startServer(18083)

	// 为了保持主函数运行
	select {}
}

/*
启动之后，postman去调用8080端口，一次次看，可以看到请求是被负载均衡器轮训分发到不同的后端服务器了
*/
