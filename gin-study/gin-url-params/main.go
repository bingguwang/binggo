package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//当请求带参数时,简单参数

func main() {
	/*
		获取get请求参数:
			Query,DefaultQuery,GetQuery
		获取post请求参数:
			postForm,DefaultpostForm,GetpostForm
		获取URL路径参数：/user/:id这种形式的路径
			c.Param("id")

	*/

	r := gin.Default()

	r.GET("/game", func(c *gin.Context) { //也就是?后用=&拼接的那种
		name, ok := c.GetQuery("name")
		if ok {
			fmt.Println(name)
		}
	})

	r.GET("/user/:id/*action", func(c *gin.Context) { //输入请求 /user/1/a/b会得到 1,和 a/b
		ac := c.Param("action")
		id := c.Param("id")
		fmt.Println(id)
		fmt.Println(ac)
	})

	r.Run(":8088")
}
