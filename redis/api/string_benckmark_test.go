package api

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

// 检测set命令
func BenchmarkRedisProxyer_Set(b *testing.B) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		b.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testValue := "value"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = proxy.Set(testKey, testValue, 1*time.Second)
		if err != nil {
			b.Fatal(err)
		}
	}

}

// 测试get函数
func BenchmarkRedisProxyer_Get(b *testing.B) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		b.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	var testString = "test"

	err = proxy.Set(testKey, testString, 0)
	if err != nil {
		b.Fatal(err)
	}
	defer proxy.Del(testKey)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proxy.Get(testKey)
		if err != nil {
			b.Fatal("读取string", err)
		}
	}
}

// 测试原生的get函数
func BenchmarkRedisProxyer_Get1(b *testing.B) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		b.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	var testString = "test"

	err = proxy.Set(testKey, testString, 0)
	if err != nil {
		b.Fatal(err)
	}
	defer proxy.Del(testKey)

	db := redis.NewClient(&redis.Options{
		Addr:         "192.168.3.64:6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	var ctx = context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := db.Get(ctx, testKey).Result()
		if err != nil {
			b.Fatal("读取string", err)
		}
	}
}

// 测试pipeline
func BenchmarkRedisProxyer_Pipeline(b *testing.B) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		b.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	var testString = "test"

	err = proxy.Set(testKey, testString, 0)
	if err != nil {
		b.Fatal(err)
	}
	defer proxy.Del(testKey)

	// 相当于 QPS * 10
	args := make([][]interface{}, 0)
	for i := 0; i < 10; i++ {
		args = append(args, []interface{}{"get", testKey})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proxy.Pipeline(args)
		if err != nil {
			b.Fatal("读取string", err)
		}
	}
}
