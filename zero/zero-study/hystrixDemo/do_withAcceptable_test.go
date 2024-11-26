package main

import (
	"github.com/zeromicro/go-zero/core/breaker"
	"testing"
)

/**
DoWithAcceptable 支持自定义的采集指标，
可以自主控制哪些情况是可以接受，哪些情况是需要加入熔断指标采集窗口的。
*/

// DoWithAcceptable第三个参数是Acceptable，就是判断错误是否可接受
func TestDoWithAcceptable(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if err := breaker.DoWithAcceptable("test", func() error {
			return mockRequest()
		}, func(err error) bool { // 当 mock 的http 状态码部位500时都会被认为是正常的，否则加入错误窗口
			me, ok := err.(mockError)
			if ok {
				return me.status != 500 // 状态码部位500时都会被认为是正常，不会被判定调用失败，不会计入熔断指标里
			}
			return false
		}); err != nil {
			println(err.Error())
		}
	}
}
