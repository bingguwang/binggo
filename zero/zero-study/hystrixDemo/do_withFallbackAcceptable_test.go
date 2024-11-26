package main

import (
	"errors"
	"github.com/zeromicro/go-zero/core/breaker"
	"testing"
)

// DoWithFallbackAcceptable 支持采集指标自定义，也支持熔断回调。
func TestFallbackAcceptable(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if err := breaker.DoWithFallbackAcceptable("test", func() error {
			return mockRequest()
		}, func(err error) error {
			//发生了熔断，这里可以自定义熔断错误转换
			return errors.New("当前服务不可用，请稍后再试")
		}, func(err error) bool { // 当 mock 的http 状态码部位500时都会被认为是正常的，否则加入错误窗口
			me, ok := err.(mockError)
			if ok {
				return me.status != 500
			}
			return false
		}); err != nil {
			println(err.Error())
		}
	}
}
