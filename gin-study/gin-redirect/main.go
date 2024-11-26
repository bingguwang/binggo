package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/a", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/b") //会重定向到/b
	})

	r.GET("/b", func(c *gin.Context) {
		c.String(http.StatusOK, "bbbb")
	})

	r.Run(":8088")
}
