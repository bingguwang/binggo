package main

//当请求带参数时,与结构体相关的参数，或者是文件参数
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/a", func(c *gin.Context) {
		req := make(map[string]interface{}, 0)
		if err := c.ShouldBind(&req); err != nil {
			return
		}
		fmt.Println("req---", req)
		return
	})
	// 底层用的都是 ShouldBindWith

	/**
	ShouldBindJSON和ShouldBind的区别：
	ShouldBindJSON只能解析 content-type是ShouldBindJSON的请求
	ShouldBind可以自动选择合适的解析器，application/json、application/x-www-form-urlencoded、multipart/form-data 等


	ShouldBind和Bind的区别：
	二者都会自动选择合适的解析器
	Bind不需要错误处理，会自动返回错误，不需要显示处理错误
	ShouldBind需要自己处理错误，但这样更灵活

	BindJSON 和Bind的区别：
	前者只能解析 content-type是ShouldBindJSON的请求，后者会自动选择
	*/

	//与结构体绑定的形式的参数
	r.POST("/singer", func(c *gin.Context) {
		param := c.Param("id")
		fmt.Println(param)
		s := Singer{}
		b := make(map[string]interface{})
		if err := c.BindJSON(&b); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(b)

		// todo 特别注意的，参数解析只能进行一次，解析完再次解析的时候会EOF，比如这里BindJson已经解析过了，下面再解析都会EOF

		if err := c.ShouldBind(&s); err == nil { //用ShouldBind接收结构体绑定的形式的参数
			fmt.Println(s.Age, " ", s.Name)
		} else {
			// ShouldBind需要自己处理错误
			fmt.Println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		if err := c.Bind(&s); err != nil {
			// Bind 已经自动处理错误响应，所以这里不需要再次处理
			fmt.Println(err.Error())
			return
		} else {
			fmt.Println(s.Age, " ", s.Name)
		}
	})

	//表单数据类型
	r.POST("/login", func(c *gin.Context) {
		var a Login
		if err := c.Bind(&a); err != nil { //bind()会根据请求头中content-type自动推断，解析并绑定form格式
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//校验
		if a.Password != "123" || a.UserName != "wb" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304", "msg": "登录失败，密码或账号不对"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200", "msg": "登录成功"})
	})

	//文件类型参数
	r.POST("/upload", func(c *gin.Context) {
		// file, err := c.FormFile("file") //FormFile方法获取文件参数,但是没有FileHeader头，只有multipart.File
		_, headers, err := c.Request.FormFile("file") //(multipart.File, *multipart.FileHeader, error)
		if err != nil {
			c.String(500, "上传图片出错")
		}
		if headers.Size > 1024*1024*30 { //大于30MB
			c.String(http.StatusForbidden, "文件不能大于30MB")
			return
		}
		fmt.Println(headers.Header)
		if headers.Header.Get("Content-type") != "audio/mpeg" {
			c.String(http.StatusForbidden, "只能上传音频文件")
			return
		}
		c.SaveUploadedFile(headers, "xxxx") //保存文件,第二个参数是保存文件路径
		c.String(http.StatusOK, "xxxx")
	})
	r.Run(":8881")

	//TODO 上传多个文件的在中间件案例中
}

type Singer struct {
	Id   int    `json:"id"`
	Age  int    `json:"age" form:"age"`
	Name string `json:"name" form:"name"`
}

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	UserName string `form:"userName" json:"userName" xml:"userName" uri:"user" binding:"required"` //名字要和表单对应
	Password string `form:"password" json:"password" xml:"password" uri:"password" binding:"required"`
}
