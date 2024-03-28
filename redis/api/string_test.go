package api

import (
	"fmt"
	"testing"
	"time"
)

var minIdleConns = 1

// 测试检查键值存在
func TestRedisProxyer_Exists(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "not exist key"

	ret, err := proxy.Exists(testKey)
	if err != nil {
		t.Fatal(err)
	}
	if ret == 1 {
		t.Fatal("检查有误")
	}
}

// 检测set命令和超时
func TestRedisProxyer_Set(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testValue := "value"

	err = proxy.Set(testKey, testValue, 1*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := proxy.Exists(testKey)
	if ret == 0 {
		t.Fatal("写入失败")
	}

	time.Sleep(2 * time.Second)

	ret, err = proxy.Exists(testKey)
	if ret == 1 {
		t.Fatal("过期删除失败")
	}
}

// 检测del删除单个key
func TestRedisProxyer_Del(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testValue := "value"

	err = proxy.Set(testKey, testValue, 0)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := proxy.Exists(testKey)
	if ret == 0 {
		t.Fatal("写入失败")
	}

	err = proxy.Del(testKey)
	if err != nil {
		t.Fatal("删除失败", err)
	}

	ret, err = proxy.Exists(testKey)
	if ret == 1 {
		t.Fatal("删除检查失败")
	}
}

// 测试一系列get函数
func TestRedisProxyer_Get(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	var testString = "test"

	err = proxy.Set(testKey, testString, 0)
	if err != nil {
		t.Fatal(err)
	}

	value, err := proxy.Get(testKey)
	if err != nil {
		t.Fatal("读取string", err)
	}
	if value != testString {
		t.Fatal("读取string值错误")
	}

	var testBytes = []byte("test")
	err = proxy.Set(testKey, testBytes, 0)
	if err != nil {
		t.Fatal(err)
	}
	_, err = proxy.GetBytes(testKey)
	if err != nil {
		t.Fatal("读取byte", err)
	}

	var testBool = true
	err = proxy.Set(testKey, testBool, 0)
	if err != nil {
		t.Fatal(err)
	}
	boolValue, err := proxy.GetBool(testKey)
	if err != nil {
		t.Fatal("读取bool", err)
	}
	if boolValue != testBool {
		t.Fatal("读取bool值错误")
	}

	var testFloat64 float64 = 1.1
	err = proxy.Set(testKey, testFloat64, 0)
	if err != nil {
		t.Fatal(err)
	}
	Float64Value, err := proxy.GetFloat64(testKey)
	if err != nil {
		t.Fatal("读取float64", err)
	}
	if Float64Value != testFloat64 {
		t.Fatal("读取float64值错误")
	}

	var testFloat32 float32 = 1.2
	err = proxy.Set(testKey, testFloat32, 0)
	if err != nil {
		t.Fatal(err)
	}
	Float32Value, err := proxy.GetFloat32(testKey)
	if err != nil {
		t.Fatal("读取float32", err)
	}
	if Float32Value != testFloat32 {
		t.Fatal("读取float32值错误")
	}

	var testInt int = 1
	err = proxy.Set(testKey, testInt, 0)
	if err != nil {
		t.Fatal(err)
	}
	IntValue, err := proxy.GetInt(testKey)
	if err != nil {
		t.Fatal("读取int", err)
	}
	if IntValue != testInt {
		t.Fatal("读取int值错误")
	}

	var testInt64 int64 = 1
	err = proxy.Set(testKey, testInt64, 0)
	if err != nil {
		t.Fatal(err)
	}
	Int64Value, err := proxy.GetInt64(testKey)
	if err != nil {
		t.Fatal("读取int64", err)
	}
	if Int64Value != testInt64 {
		t.Fatal("读取int64值错误")
	}

	var testUint64 uint64 = 2
	err = proxy.Set(testKey, testUint64, 0)
	if err != nil {
		t.Fatal(err)
	}
	Uint64Value, err := proxy.GetUint64(testKey)
	if err != nil {
		t.Fatal("读取uint64", err)
	}
	if Uint64Value != testUint64 {
		t.Fatal("读取uint64值错误")
	}
	_ = proxy.Del(testKey)
}

// 测试mset
func TestRedisProxyer_MSet(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey1 := "test_key1"
	testValue1 := "value1"
	testKey2 := "test_key2"
	testValue2 := "value2"

	err = proxy.MSet(testKey1, testValue1, testKey2, testValue2)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey1, testKey2)

	value, err := proxy.Get(testKey1)
	if err != nil {
		t.Fatal("读取", err)
	}
	if value != testValue1 {
		t.Fatal("检查值1错误")
	}

	value, err = proxy.Get(testKey2)
	if err != nil {
		t.Fatal("读取", err)
	}
	if value != testValue2 {
		t.Fatal("检查值2错误")
	}
}

func TestRedisProxyer_MGet(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey1 := "test_key1"
	testValue1 := "value1"
	testKey2 := "test_key2"
	testValue2 := "value2"

	err = proxy.MSet(testKey1, testValue1, testKey2, testValue2)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey1, testKey2)

	values, err := proxy.MGet(testKey1, testKey2, "not_exist_key")
	if err != nil {
		t.Fatal("读取", err)
	}
	if values[0].(string) != testValue1 {
		t.Fatal("检查值1错误")
	}

	if values[1].(string) != testValue2 {
		t.Fatal("检查值2错误")
	}

	if values[2] != nil {
		t.Fatal("检查值3错误")
	}
}

// 测试单点故障切换
func TestRedisProxyer_SinglePointFailure(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testValue := "value"

	for count := 60; count > 0; count -= 1 {
		t.Logf("=== count: %d ===\n", count)
		err = proxy.Set(testKey, fmt.Sprintf("%s%d", testValue, count), 0)
		if err != nil {
			time.Sleep(time.Second * 5)
			t.Log(err)
			continue
		}

		value, err := proxy.Get(testKey)
		if err != nil {
			t.Log("读取错误", err)
			continue
		}
		t.Logf("get value: %s\n", value)

		err = proxy.Del(testKey)
		if err != nil {
			t.Fatal("删除错误", err)
		}

		time.Sleep(time.Second)
	}
}
