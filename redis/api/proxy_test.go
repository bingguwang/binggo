package api

import (
	"testing"
)

// 测试是否可以连接redis
func TestNewRedisProxy(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()
}
