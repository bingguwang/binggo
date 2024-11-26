package main

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/breaker"
	"math/rand"
	"testing"
	"time"
)

type mockError struct {
	status int
}

func (e mockError) Error() string {
	return fmt.Sprintf("HTTP STATUS: %d", e.status)
}

/**
Do 方法是不能自定义熔断器的参数的
*/

func TestDo(t *testing.T) {
	for i := 0; i < 1000; i++ {
		// 会以传入的test为key创建一个breaker实例
		if err := breaker.Do("test", func() error {
			// 调用mockRequest方法，这里可以是各种调用，比如rpc调用、对mysql redis等中间件的调用
			// 因为熔断器是一种客户端的机制，为的就是客户端停止'过分的调用'
			return mockRequest()
		}); err != nil {
			// 没有回调，只能得到调用的函数所返回的错误
			println("ERR：", err.Error())
		}
	}
}

func mockRequest() error {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num := r.Intn(100)
	if num%4 == 0 {
		return nil
	} else if num%5 == 0 {
		return mockError{status: 500}
	}
	return errors.New("dummy")
}
