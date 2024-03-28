package api

import (
	"testing"
)

func TestRedisProxyer_ZAdd(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testValue := map[interface{}]float64{
		25.1: 202008030905,
		25.3: 202008030906,
		25.9: 202008030907,
	}

	err = proxy.ZAdd(testKey, testValue)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)

	ret, err := proxy.ZRangeByScore(testKey, "202008030905", "202008030907")
	t.Log(ret)
}

func TestRedisProxyer_ZRem(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testValue := map[interface{}]float64{
		25.1: 202008030905,
		25.3: 202008030906,
		25.9: 202008030907,
	}

	err = proxy.ZAdd(testKey, testValue)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)

	err = proxy.ZRem(testKey, 25.3)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := proxy.ZRangeByScore(testKey, "202008030905", "202008030907")
	t.Log(ret)
}

func TestRedisProxyer_ZRemRangeByScore(t *testing.T) {
	proxy, err := NewRedisProxy(minIdleConns)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Close()

	testKey := "test_key"
	testValue := map[interface{}]float64{
		25.1: 202008030905,
		25.3: 202008030906,
		25.9: 202008030907,
		25.7: 202008030908,
	}

	err = proxy.ZAdd(testKey, testValue)
	if err != nil {
		t.Fatal(err)
	}
	defer proxy.Del(testKey)

	ret, err := proxy.ZRangeByScore(testKey, "202008030905", "202008030908")
	t.Log(ret)

	err = proxy.ZRemRangeByScore(testKey, "202008030905", "202008030907")
	if err != nil {
		t.Fatal(err)
	}

	ret, err = proxy.ZRangeByScore(testKey, "202008030905", "202008030908")
	t.Log(ret)
}
