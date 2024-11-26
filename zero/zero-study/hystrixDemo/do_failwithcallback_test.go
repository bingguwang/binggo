package main

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/breaker"
	"testing"
)

// DoWithFallback 默认采用错误率来判断服务是否可用，不支持指标自定义，但是支持熔断回调。
// 第三个参数是回调函数，在发生熔断时，请求被的
func TestDoWithFallback(t *testing.T) {

	for i := 0; i < 1000; i++ {
		if err := breaker.DoWithFallback("test", func() error {
			return mockRequest()
		}, func(err error) error {
			// 发生了熔断，这里可以自定义熔断错误转换
			fmt.Println("发生了熔断，这里可以自定义熔断错误转换")
			return errors.New("当前服务不可用，请稍后再试")
		}); err != nil {
			println(err.Error())
		}
	}
}
