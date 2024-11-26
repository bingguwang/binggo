package main

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// Fields type, used to pass to `WithFields`.
type Fields = logrus.Fields

// Debug log
var Debug func(args ...interface{})

// Info log
var Info func(args ...interface{})

// Warn log
var Warn func(args ...interface{})

// Error log
var Error func(args ...interface{})

// Fatal will Exit(1) after logging
var Fatal func(args ...interface{})

// WithFields Adds a struct of fields to the log entry
var WithFields func(fields Fields) *logrus.Entry

var (
	logonce sync.Once
	LogFile *lumberjack.Logger
)

func SetupLogger(logPath string) *logrus.Logger {

	logger := logrus.New()
	writer, err := rotatelogs.New(
		logPath+".%Y%m%dT%H%M%S.log",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logPath),

		// WithRotationTime设置日志分割的时间
		rotatelogs.WithRotationTime(30*time.Second),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
		rotatelogs.WithMaxAge(time.Duration(3000000)*time.Minute),
		//rotatelogs.WithRotationCount(3),
	)
	if err == nil {
		fmt.Println("用rotatelogs 切割")
		logger.SetOutput(writer)
	} else {
		fmt.Println("使用rotatelogs 切割失败，转用稍笨拙一点的lumberjack切割")
		if LogFile == nil {
			logonce.Do(func() {
				LogFile = &lumberjack.Logger{
					Filename:   logPath, // 指定日志文件路径
					MaxSize:    1,       // 每个日志文件最大尺寸（MB）
					MaxBackups: 30,      // 最多保留旧日志文件的个数
					MaxAge:     7,       // 保留最近 N 天的日志文件
					Compress:   false,   // 是否压缩旧的日志文件
					LocalTime:  true,    // 设置 LocalTime 为 true，使时间本地化
				}
			})
		}
		logger.SetOutput(LogFile)
	}
	logger.SetLevel(logrus.InfoLevel)
	//设置日志格式
	/*logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})*/
	Debug = logger.Debug
	Info = logger.Info
	Warn = logger.Warn
	Error = logger.Error
	Fatal = logger.Fatal

	WithFields = logger.WithFields
	return logger
}
