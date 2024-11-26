package main

import (
	// 导入session包
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	// 导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
	// 导入gin框架包
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//创建一个基于cookie的存储引擎，password123456是加密的密钥
	store := cookie.NewStore([]byte("password123456"))

	//设置session中间件
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		//初始化session对象
		session := sessions.Default(c)
		if session.Get("name") != "wb" {
			session.Set("name", "wb")
			session.Save()
			//删除整个session
			// session.Clear()
		}
		c.JSON(http.StatusOK, gin.H{"name": session.Get("name")})
		fmt.Printf("%#v\n", store)
	})

	r.Run(":8088")

}

//TODO     "github.com/gorilla/sessions"也可以用于session
