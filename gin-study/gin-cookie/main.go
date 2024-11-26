package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//gin的cookie和http中的几乎一样，只是封装了一下

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("mycookie")
		if err != nil { //不存在cookie则创建
			cookie = "AE86"
			// 给客户端设置cookie
			//  maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			//   secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("mycookie", "AE86", 60, "/", "localhost", false, true)
		}
		fmt.Println("cookie的值是：", cookie)
	})

	r.Run(":8088")
}
