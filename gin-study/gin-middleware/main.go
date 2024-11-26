package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*
	》》》》》》》》	中间件
	请求-------响应这段生命周期内，可以注册多个中间件，每个中间件执行不同的功能
	gin.Default其实就默认使用了2个中间件Logger()//打印请日志, Recovery()//拦截panic错误
	全局中间件：对所有的请求生效

*/
//全局中间件。用use方法注册的
func main() {
	r := gin.Default()

	//注册中间件,全局中间件
	r.Use(Middleware()) //(engine *Engine) Use(middleware ...HandlerFunc)

	r.GET("/a", func(c *gin.Context) {
		req, _ := c.Get("req")
		fmt.Println(req)
		c.JSON(200, gin.H{"req": req}) //其实可以看到gin.Context的全局性，即使是在了方法中的形参，但是是全局共享的，可以在别的地方修改它
	})

	//注册中间件，局部中间件，只作用于局部,可以看到请求/a就没有这个中间件的效果,/a只有全局中间件的效果，但全局中间件的效果/b也有
	r.GET("/b", Middleware_2(), func(c *gin.Context) {
		req, _ := c.Get("req")
		fmt.Println(req)
		c.JSON(200, gin.H{"req": req})
		resp, _ := c.Get("resp")
		fmt.Println(resp)
		c.JSON(200, gin.H{"resp": resp})
	})

	//GitHub有很多中间件可以用

	r.Run(":8088")
}

//自定义一个中间件
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("开始执行中间件")
		c.Set("req", "中间件")         //在上下文中设置一个k-v,假设是中间件做的工作
		c.Next()                    //执行下一个中间件，如果不写，下一个中间件就执行不了，所以最好每个中间件中都加c.next
		status := c.Writer.Status() //查询请求状态码
		fmt.Println("中间件结束执行", status)
		t2 := time.Since(t)
		fmt.Println("执行时间：", t2)
	}
}

func Middleware_2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("resp", "这个中间件只作用于局部的")
		c.Next()
	}
}
