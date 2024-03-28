package api

import (
	"testing"
)

// 检查 SAdd 和 SIsMember 命令
func TestRedisProxyer_SAdd(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testValue := []string{"1", "2"}

	err = proxy.SAdd(testKey, testValue)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)

	//t.Log(proxy.SMembers(testKey))
	for _, value := range testValue {
		isMember, _ := proxy.SIsMember(testKey, value)
		if !isMember {
			t.Fatal("检查成员错误")
		}
	}
}

// 移除元素
func TestRedisProxyer_SRem(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testValue := []string{"1", "2"}

	err = proxy.SAdd(testKey, testValue)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)

	err = proxy.SRem(testKey, "2")
	if err != nil {
		t.Fatal(err)
	}
	// ["1"]
	t.Log(proxy.SMembers(testKey))
}

// 差集
func TestRedisProxyer_SDiff(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey1 := "test_key1"
	testValue1 := []string{"1", "2"}
	testKey2 := "test_key2"
	testValue2 := []string{"3", "2"}

	err = proxy.SAdd(testKey1, testValue1)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey1)
	err = proxy.SAdd(testKey2, testValue2)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey2)

	diffSet, err := proxy.SDiff(testKey1, testKey2)
	if err != nil {
		t.Fatal(err)
	}
	// ["1"]
	t.Log(diffSet)
}

// 交集
func TestRedisProxyer_SInter(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey1 := "test_key1"
	testValue1 := []string{"1", "2"}
	testKey2 := "test_key2"
	testValue2 := []string{"3", "2"}

	err = proxy.SAdd(testKey1, testValue1)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey1)
	err = proxy.SAdd(testKey2, testValue2)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey2)

	interSet, err := proxy.SInter(testKey1, testKey2)
	if err != nil {
		t.Fatal(err)
	}
	// ["2"]
	t.Log(interSet)
}

// 并集
func TestRedisProxyer_SUnion(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey1 := "test_key1"
	testValue1 := []string{"1", "2"}
	testKey2 := "test_key2"
	testValue2 := []string{"3", "2"}

	err = proxy.SAdd(testKey1, testValue1)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey1)
	err = proxy.SAdd(testKey2, testValue2)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey2)

	unionSet, err := proxy.SUnion(testKey1, testKey2)
	if err != nil {
		t.Fatal(err)
	}
	// ["1", "2", "3"]
	t.Log(unionSet)
}
