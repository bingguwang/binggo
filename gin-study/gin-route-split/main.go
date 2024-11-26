package main

import (
	"github.com/wbing441282413/ginTest/gin-route-split/route/order"
	"github.com/wbing441282413/ginTest/gin-route-split/route/user"
	"github.com/wbing441282413/ginTest/gin-route-split/routers"
)

//》》》》》》》》》》》》》    路由拆分
func main() {

	//项目简单的时候会这样路由，但是路由多的时候应该不要放在main中，
	// 应该把路由部分的代码都拆分出来，形成一个单独的文件或包
	// r := gin.Default()
	// r.GET("/a", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "wbing",
	// 	})
	// })

	//加载定义的路由
	routers.Include(order.OrderGet, user.UserGet)
	//初始化路由，绑定路由与处理函数
	r := routers.Init()
	r.Run(":8088")
}
