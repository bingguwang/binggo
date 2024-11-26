package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/u1") //返回*RouterGroup，它有成员Engine，所以可以调用get等方法
	{
		v1.GET("/", getU1)
		v1.GET("/orders", getOrders)
		v1.GET("/refunds", getRefunds)
	}

	v2 := r.Group("/")
	{
		v2.GET("/", getIndex)
		v2.GET("/haha", getHaha)
		v2.POST("/post", func(c *gin.Context) { //一组内可以是不同的请求
			c.String(200, "post")
		})
	}
	r.Run(":8088")
}
func getOrders(c *gin.Context) {
	c.String(200, "orders")
}
func getRefunds(c *gin.Context) {
	c.String(200, "refunds")
}
func getU1(c *gin.Context) {
	c.String(200, "u1")
}

func getIndex(c *gin.Context) {
	c.String(200, "index")
}
func getHaha(c *gin.Context) {
	c.String(200, "haha")
}
