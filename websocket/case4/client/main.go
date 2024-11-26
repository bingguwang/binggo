package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var countRecv int

func main() {
	f := func() *gin.Engine {
		r := gin.Default()
		r.POST("/collect/message", func(c *gin.Context) {
			mp := make(map[string]interface{})
			if err := c.BindJSON(&mp); err != nil {
				fmt.Println("[error] ", err.Error())
				return
			}
			fmt.Println("received: ", mp)
			countRecv++
			fmt.Println("共收到消息:", countRecv, "条")
		})
		r.GET("/gb/keepAlive", func(c *gin.Context) {
			c.JSON(http.StatusOK, "OK")
		})
		return r
	}

	httpSrv := &http.Server{
		Addr:    ":7799",
		Handler: f(),
	}
	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err.Error())
		return
	}
}
