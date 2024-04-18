package main

import (
	"binggo/aboutLog/sirupsenlogrus/log"
	"binggo/aboutLog/sirupsenlogrus/utils"
	"time"
)

var (
	loger         log.Logger
	logerShowMore log.Logger
	debug         bool
)

// 可能需要先开启开发者模式
func main() {
	loger = utils.NewLogrusLogger(log.InfoLevel, "我是打印输出时加在日志里的前缀", nil, nil, newLfsHook(true, logPathPrefix))
	var count = 0

	for {
		time.Sleep(1 * time.Second)
		count++
		loger.Info("time:", time.Now(), "序号:", count)
	}
}
