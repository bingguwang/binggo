package utils

import (
	"binggo/aboutLog/sirupsenlogrus/log"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type MyLogger struct {
	Logger *log.LogrusLogger
	level  log.Level
}

func (ml *MyLogger) Level() string {
	switch ml.level {
	case log.PanicLevel:
		return "Panic"
	case log.FatalLevel:
		return "Fatal"
	case log.ErrorLevel:
		return "Error"
	case log.WarnLevel:
		return "Warn"
	case log.InfoLevel:
		return "Info"
	case log.DebugLevel:
		return "Debug"
	case log.TraceLevel:
		return "Trace"
	}
	return "Unkown"
}

var (
	loggers map[string]*MyLogger
)

func init() {
	loggers = make(map[string]*MyLogger)
}

func NewLogrusLogger(level log.Level, prefix string, fields log.Fields, writer io.Writer, hook logrus.Hook) log.Logger {
	if logger, found := loggers[prefix]; found {
		return logger.Logger.WithPrefix(prefix)
	}
	l := logrus.New()
	if writer != nil {
		l.SetOutput(writer)
	}
	if hook != nil {
		l.AddHook(hook)
	}
	l.Level = logrus.ErrorLevel
	l.Formatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		ForceColors:     true,
		ForceFormatting: true,
		DisableColors:   true,
	}
	//l.SetReportCaller(true) // 见 main/logInit.go
	logger := log.NewLogrusLogger(l, "main", fields)
	loggers[prefix] = &MyLogger{
		Logger: logger,
		level:  level,
	}
	logger.SetLevel(level)
	return logger.WithPrefix(prefix)
}

func SetLogLevel(prefix string, level log.Level) error {
	if logger, found := loggers[prefix]; found {
		logger.level = level
		logger.Logger.SetLevel(level)
		return nil
	}
	return fmt.Errorf("logger [%v] not found", prefix)
}

func GetLoggers() map[string]*MyLogger {
	return loggers
}
