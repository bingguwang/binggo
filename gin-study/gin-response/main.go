package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Gin的hTTP请求的响应
// gin.Context上下文支持多种返回处理结果，字符串，JSON，XML等格式
func main() {
	r := gin.Default()
	s := &Singer{
		Name: "dt",
		Age:  52,
	}

	r.GET("/string", func(c *gin.Context) {
		// Headers里需要设置值，可以使用Header ,k-v string类型的键值对
		//c.Header("name", "wbing") //value得是string
		//c.Header("status", "200")
		//c.Header("token", "SADOIHSABWQJBJ2312IU321H")
		OkWithMsg(c, fmt.Sprintf("hello %s ,你好帅", "wbing"))
		return
	})

	// msg最好是const，预先定义好,比如message.go里的那种
	r.GET("/json", func(c *gin.Context) {
		OkWithData(c, "OK", s) //传入结构体的指针,响应结果为{"Name":"dt","Age":52}
		return
	})

	// ----------------------- 下面是一些不常用的特殊的返回格式 ------------------------
	//XML格式
	r.GET("/xml", func(c *gin.Context) {

		//c.XML(200, s) //xml的根节点的名字默认是结构体的名字
		OkWithXmlData(c, "OK", s) //xml的根节点的名字默认是结构体的名字
	})

	//文件格式
	r.GET("/file", func(c *gin.Context) {
		// c.File("C:/Users/WBing/Pictures/rmrf.gif") //参数是文件地址
		c.File("./rmrf.gif") //参数是文件地址
	})

	// 4.YAML响应
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan"})
	})

	r.Run(":8881")
}

type Singer struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}
