package api

import (
	"testing"
)

// 测试检查hash键值存在
func TestRedisProxyer_HExists(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "not exist key"
	testField := "field"

	ret, err := proxy.HExists(testKey, testField)
	if err != nil {
		t.Fatal(err)
	}
	if ret {
		t.Fatal("检查有误")
	}
}

// 检测hset命令和超时
func TestRedisProxyer_HSet(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testField := "field"
	testValue := "value"

	err = proxy.HSet(testKey, testField, testValue)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := proxy.HExists(testKey, testField)
	if !ret {
		t.Fatal("写入失败")
	}
}

// 检测del删除单个key
func TestRedisProxyer_HDel(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testField := "field"
	testValue := "value"

	err = proxy.HSet(testKey, testField, testValue)
	if err != nil {
		t.Fatal(err)
	}

	err = proxy.HDel(testKey, testField)
	if err != nil {
		t.Fatal("删除失败", err)
	}

	ret, err := proxy.HExists(testKey, testField)
	if ret {
		t.Fatal("删除检查失败")
	}
}

// 测试一系列get函数
func TestRedisProxyer_HGet(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testField := "field"
	var testString = "test"

	err = proxy.HSet(testKey, testField, testString)
	if err != nil {
		t.Fatal(err)
	}

	value, err := proxy.HGet(testKey, testField)
	if err != nil {
		t.Fatal("读取string", err)
	}
	if value != testString {
		t.Fatal("读取string值错误")
	}

	var testBytes = []byte("test")
	err = proxy.HSet(testKey, testField, testBytes)
	if err != nil {
		t.Fatal(err)
	}
	_, err = proxy.HGetBytes(testKey, testField)
	if err != nil {
		t.Fatal("读取byte", err)
	}

	var testBool = true
	err = proxy.HSet(testKey, testField, testBool)
	if err != nil {
		t.Fatal(err)
	}
	boolValue, err := proxy.HGetBool(testKey, testField)
	if err != nil {
		t.Fatal("读取bool", err)
	}
	if boolValue != testBool {
		t.Fatal("读取bool值错误")
	}

	var testFloat64 float64 = 1.1
	err = proxy.HSet(testKey, testField, testFloat64)
	if err != nil {
		t.Fatal(err)
	}
	Float64Value, err := proxy.HGetFloat64(testKey, testField)
	if err != nil {
		t.Fatal("读取float64", err)
	}
	if Float64Value != testFloat64 {
		t.Fatal("读取float64值错误")
	}

	var testFloat32 float32 = 1.2
	err = proxy.HSet(testKey, testField, testFloat32)
	if err != nil {
		t.Fatal(err)
	}
	Float32Value, err := proxy.HGetFloat32(testKey, testField)
	if err != nil {
		t.Fatal("读取float32", err)
	}
	if Float32Value != testFloat32 {
		t.Fatal("读取float32值错误")
	}

	var testInt int = 1
	err = proxy.HSet(testKey, testField, testInt)
	if err != nil {
		t.Fatal(err)
	}
	IntValue, err := proxy.HGetInt(testKey, testField)
	if err != nil {
		t.Fatal("读取int", err)
	}
	if IntValue != testInt {
		t.Fatal("读取int值错误")
	}

	var testInt64 int64 = 1
	err = proxy.HSet(testKey, testField, testInt64)
	if err != nil {
		t.Fatal(err)
	}
	Int64Value, err := proxy.HGetInt64(testKey, testField)
	if err != nil {
		t.Fatal("读取int64", err)
	}
	if Int64Value != testInt64 {
		t.Fatal("读取int64值错误")
	}

	var testUint64 uint64 = 2
	err = proxy.HSet(testKey, testField, testUint64)
	if err != nil {
		t.Fatal(err)
	}
	Uint64Value, err := proxy.HGetUint64(testKey, testField)
	if err != nil {
		t.Fatal("读取uint64", err)
	}
	if Uint64Value != testUint64 {
		t.Fatal("读取uint64值错误")
	}
	_ = proxy.Del(testKey)
}
