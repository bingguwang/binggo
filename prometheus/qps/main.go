package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

// 只可增加的一个计数器
var req_counter_vec = prometheus.NewCounterVec( // 定义了一恶搞counter vector
	prometheus.CounterOpts{
		Name: "req_counter_vec",
		Help: "request counter vector",
	},
	[]string{"endpoint"}, // 以endpoint区分counter vector里不同的元素
)

func main() {
	log.SetOutput(os.Stdout)
	prometheus.MustRegister(req_counter_vec)

	http.Handle("/metrics", promhttp.Handler()) // 这个接口提供了一组关于应用程序或服务运行状况、性能和资源使用的指标数据。被Prometheus抓取
	http.HandleFunc("/hello", HelloHandler)     // 业务接口，以及业务接口的处理逻辑handler

	log.Println("运行成功:20001")
	errChan := make(chan error)
	go func() {
		// 业务服务在20001端口运行，
		fmt.Println("运行成功:20001")
		log.Println("运行成功:20001")
		errChan <- http.ListenAndServe(":20001", nil)
	}()
	err := <-errChan
	if err != nil {
		fmt.Println("Hello server stop running.")
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	req_counter_vec.WithLabelValues(path).Inc()
	w.Write([]byte("调用成功"))
	time.Sleep(100 * time.Millisecond)
}
