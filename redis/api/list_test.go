package api

import (
	"testing"
	"time"
)

// 测试lpush,lrange,llen
func TestRedisProxyer_LPush(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_list"
	testValue := []string{"1", "2", "3"}
	err = proxy.LPush(testKey, testValue[0])
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)
	length, err := proxy.LLen(testKey)
	if length != 1 {
		t.Fatal("列表长度错误")
	}

	err = proxy.LPush(testKey, testValue[1:])
	if err != nil {
		t.Fatal(err)
	}
	length, _ = proxy.LLen(testKey)

	ret, err := proxy.LRange(testKey, 0, length-1)
	if err != nil {
		t.Fatal(err)
	}
	for idx, v := range testValue {
		if v != ret[int(length-1)-idx] {
			t.Fatal("列表元素错误")
		}
	}
}

// 测试RPush， LPop
func TestRedisProxyer_LPop(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_list"
	testValue := []string{"1", "2", "3"}
	err = proxy.RPush(testKey, testValue)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)
	length, _ := proxy.LLen(testKey)

	for idx := 0; idx < int(length); idx++ {
		value, err := proxy.LPop(testKey)
		if err != nil {
			t.Fatal(err)
		}
		if value != testValue[idx] {
			t.Fatal("LPop列表元素错误")
		}
	}

}

// 测试LPopInt
func TestRedisProxyer_LPopInt(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_list"
	testValue := []int{1, 2, 3}
	// 这里不能直接传 []int 类型
	err = proxy.RPush(testKey, 1, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)
	length, _ := proxy.LLen(testKey)

	for idx := 0; idx < int(length); idx++ {
		value, err := proxy.LPopInt(testKey)
		if err != nil {
			t.Fatal(err)
		}
		if value != testValue[idx] {
			t.Fatal("LPopInt列表元素错误")
		}
	}

}

// 测试RPop
func TestRedisProxyer_RPop(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_list"
	testValue := []string{"1", "2", "3"}
	err = proxy.LPush(testKey, testValue)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)
	length, _ := proxy.LLen(testKey)

	for idx := 0; idx < int(length); idx++ {
		value, err := proxy.RPop(testKey)
		if err != nil {
			t.Fatal(err)
		}
		if value != testValue[idx] {
			t.Fatal("RPop列表元素错误")
		}
	}
}

// 测试RPopInt
func TestRedisProxyer_RPopInt(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_list"
	testValue := []int{1, 2, 3}
	// 这里不能直接传 []int 类型
	err = proxy.LPush(testKey, 1, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)
	length, _ := proxy.LLen(testKey)

	for idx := 0; idx < int(length); idx++ {
		value, err := proxy.RPopInt(testKey)
		if err != nil {
			t.Fatal(err)
		}
		if value != testValue[idx] {
			t.Fatal("RPopInt列表元素错误")
		}
	}
}

// 测试BLPop
func TestRedisProxyer_BLPop(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_list"
	testValue := []string{"1", "2", "3"}
	err = proxy.RPush(testKey, testValue)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)
	length, _ := proxy.LLen(testKey)

	for idx := 0; idx < int(length); idx++ {
		value, err := proxy.BLPop(1*time.Second, testKey)
		if err != nil {
			t.Fatal(err)
		}
		//t.Log(value)
		if value[1] != testValue[idx] {
			t.Fatal("BLPop列表元素错误")
		}
	}

	_, err = proxy.BLPop(1*time.Second, testKey)
	if err != nil {
		t.Log("测试BLPop超时", err)
	} else {
		t.Fatal("未知异常")
	}
}

// 测试BRPop
func TestRedisProxyer_BRPop(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_list"
	testValue := []string{"1", "2", "3"}
	err = proxy.LPush(testKey, testValue)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)
	length, _ := proxy.LLen(testKey)

	for idx := 0; idx < int(length); idx++ {
		value, err := proxy.BRPop(1*time.Second, testKey)
		if err != nil {
			t.Fatal(err)
		}
		//t.Log(value)
		if value[1] != testValue[idx] {
			t.Fatal("BRPop列表元素错误")
		}
	}

	_, err = proxy.BRPop(1*time.Second, testKey)
	if err != nil {
		t.Log("测试BRPop超时", err)
	} else {
		t.Fatal("未知异常")
	}
}
