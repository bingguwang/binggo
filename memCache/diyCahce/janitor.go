package main

import "time"

// 定时清理过期
type janitor struct {
	Interval time.Duration // 清理间隔
	stop     chan bool
}

func (j *janitor) Run(c *cache) {
	for {
		select {
		case <-time.After(j.Interval):
			c.DeleteExpired()
		case <-j.stop:
			return
		}
	}
}

func stopJanitor(c *Cache) {
	c.janitor.stop <- true
}

func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{
		Interval: ci,
		stop:     make(chan bool),
	}
	c.janitor = j
	go j.Run(c)
}
