package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

/*
*
优雅的关闭gin服务
*/
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":7777",
		Handler: router,
	}

	httpPortStr := strings.Split(srv.Addr, ":")[1]
	httpPort, _ := strconv.ParseInt(httpPortStr, 10, 64)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)

	go func() {
		//增加异常捕获，从而使得关闭服务时systemctl status不显示failure
		defer func() {
			if err := recover(); err != nil {
				log.Printf("http panic found:%v", err)
			}
		}()

	loop:
		for { // 当端口被占用的时候，给端口加一在重试
			if err := srv.ListenAndServe(); err != nil {
				if NetstatAno(httpPort) {
					fmt.Println("---", httpPort)
					log.Printf("端口[%v]被占用", httpPort)
					httpPort++
					srv.Addr = fmt.Sprintf(":%v", httpPort)
					fmt.Println(httpPort)
				} else {
					log.Print("Could not start listener:", err)
				}
			} else {
				break loop
			}
			select {
			case <-quit:
				log.Println("[videoGateway] ", "video gateway exit by signal")
				break loop
			case <-time.After(3 * time.Second):
			}
		}
	}()

	fmt.Println("启动gin")
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panic("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func NetstatAno(find interface{}) bool {
	var findStr string
	switch v := find.(type) {
	case string:
		findStr = v
	case int:
		findStr = strconv.FormatInt(int64(v), 10)
	case int32:
		findStr = strconv.FormatInt(int64(v), 10)
	case int64:
		findStr = strconv.FormatInt(v, 10)
	}
	cmd := exec.Command("netstat", "-ano")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %s\n", err)
		return false
	}

	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":"+findStr) && (strings.Contains(line, "LISTENING") || strings.Contains(line, "ESTABLISHED")) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading output: %s\n", err)
	}

	return false
}
