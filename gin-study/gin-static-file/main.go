package main

import (
	"github.com/gin-gonic/gin"
)

//gin访问静态文件,有点问题，用到了在看

func main() {
	r := gin.Default()
	r.StaticFile("/rmrf.gif", "./rmrf.gif")
	r.Run(":8088")
}
