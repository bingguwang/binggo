package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

/*
logrus是最常用的日志库，看下在gin里如何使用它
一般还会用到lumberjack 日志切割
*/

func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)
		fmt.Println(latencyTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

var (
	logFilePath = ""
	fileName    = ""
	router      *gin.Engine
)

func init() {
	router = gin.Default()

	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	/*
		使用了rotatelogs后，这里不需要实现创建好日志文件了
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			fmt.Println(err.Error())
		}*/
	logFileName := time.Now().Format("2006-01-02") + ".log"
	//日志文件
	fileName = path.Join(logFilePath, logFileName)
	// 记录到文件。
	logger := SetupLogger(fileName)
	router.Use(LoggerMiddleware(logger))

}

func main() {
	router.GET("/", func(context *gin.Context) {
		Info("[info]级别的日志")
	})
	router.Run(":8881")
}
