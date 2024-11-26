package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//对于处理函数中要开新协程的，新协程应该用的是上下文的副本

func main() {
	r := gin.Default()

	//异步执行
	r.GET("/async", func(c *gin.Context) {
		copyContext := c.Copy() //为上下文拷贝一个副本
		//异步处理
		go func() { //新建的协程中要用副本上下文，不要用真的上下文
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})

	r.Run(":8088")
}
