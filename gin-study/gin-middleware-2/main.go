package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Use(TimeWare())
	v := r.Group("/", TimeWare()) //中间件作用域是分组中，分组中的路由都有效果
	{
		v.GET("/a", func(c *gin.Context) {
			time.Sleep(2 * time.Second)
		})
		v.GET("/b", func(c *gin.Context) {
			time.Sleep(5 * time.Second)
		})
	}
	r.Run(":8088")
}

func TimeWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() //执行函数
		du := time.Since(start)
		fmt.Println("执行时间：", du)
	}
}
