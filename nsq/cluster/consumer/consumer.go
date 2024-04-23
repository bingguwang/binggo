package main

import (
	"log"
	"strings"
)

// nsq 内部日志
type nsqServerLogger struct {
}

func (nsl *nsqServerLogger) Output(callDepth int, s string) error {
	log.Println("nsqServerLogger", callDepth, s[:3], strings.Trim(s[3:], " "))
	return nil
}
