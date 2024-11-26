package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//还有点问题
func main() {
	r := gin.Default()
	r.GET("/home", AuthMiddleware(), TestMiddleware(), func(c *gin.Context) {
		cookie, _ := c.Cookie("my_cookie")
		fmt.Println("cookie是： ", cookie)
		c.Next()
		c.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
	})
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("my_cookie", "123", 60, "/", "localhost", false, true)
		c.String(200, "登录成功")
		c.Redirect(http.StatusMovedPermanently, "/home")
	})

	r.Run(":8088")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("my_cookie")
		if err == nil && cookie == "123" {
			fmt.Println("有cookie")
			c.Next()
			return
		}
		fmt.Println("开始重定向")
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "没有登录"})
		// c.Abort() //后面的代码不会执行
		return
	}
}

func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "c.Abort执行后，后面其他的中间件也不会执行"})
	}
}
