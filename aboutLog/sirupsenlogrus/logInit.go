package main

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/rifflock/lfshook"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

const (
	logPath       = `./videoGateway.log`
	logPathPrefix = `./videoGateway`
)

var files []string

func newLfsHook(reportCaller bool, path string) log.Hook {
	writer, err := rotatelogs.New(
		path+".%Y%m%dT%H%M%S.log",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logPath),

		// WithRotationTime设置日志分割的时间，这里设置为24小时分割一次
		//rotatelogs.WithRotationTime(24*time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
		rotatelogs.WithMaxAge(time.Duration(3000000)*time.Minute),
		//rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}
	log.SetOutput(writer) // 在这里加入分割以及symlink才有效
	defer writer.Close()
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &MyFormatter{reportCaller: reportCaller})
	//}, &log.TextFormatter{DisableColors: true})

	return lfsHook
}

type MyFormatter struct {
	reportCaller bool
}

func (m *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	if m.reportCaller {
		pcs := make([]uintptr, 25)
		depth := runtime.Callers(10, pcs)
		frames := runtime.CallersFrames(pcs[:depth])
		f, _ := frames.Next()
		fName := filepath.Base(f.File)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d %s] %s\n", timestamp, entry.Level, fName, f.Line, f.Function, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}
