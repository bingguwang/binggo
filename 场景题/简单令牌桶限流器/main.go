package main

import (
	"fmt"
	"sync"
	"time"
)

type limiter struct {
	rate     int        // 每秒生成令牌的数
	capcity  int        // 容量
	curnum   int        // 当前令牌数
	lastFill time.Time  // 上次填充的时间
	lock     sync.Mutex // 锁
}

func newlimiter(rate int, capcity int) *limiter {
	return &limiter{
		rate:     rate,
		capcity:  capcity,
		curnum:   capcity, // 初始为满
		lastFill: time.Now(),
	}
}

func (l *limiter) consume(pop int) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	// 先填充
	internal := time.Now().Sub(l.lastFill).Seconds() // 时间间隔，s
	// 应该新增的数
	toadd := int(internal) * l.rate
	if toadd > 0 {
		l.curnum = min(l.capcity, l.curnum+toadd)
		l.lastFill = time.Now()
	}

	if l.curnum > pop {
		l.curnum -= pop
		return true
	}
	return false
}
func main() {

	l := newlimiter(5, 10)
	for i := 0; i < 20; i++ {
		if l.consume(1) {
			fmt.Println("request:", i, " allowed")
		} else {
			fmt.Println("request:", i, " denied")
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// 辅助函数：取两个数的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
