package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//根据需要定义函数用来注册子route目录中定义的路由
// Init函数用来进行路由的初始化操作：

type Option func(*gin.Engine)

var options = []Option{} //定义一个数组，元素是函数 func(*gin.Engine){}

//注册route目录下定义的路由
func Include(opts ...Option) {
	fmt.Printf("%T\n", options)
	options = append(options, opts...)
}

//初始化
func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options { //每个元素是一个 func(*gin.Engine){}，把参数传入调用，就可以实现注册所有定义的路由了
		opt(r)
	}
	return r
}
