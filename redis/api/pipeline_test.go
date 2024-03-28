package api

import (
	"testing"
	"time"
)

func TestRedisProxyer_Pipelined(t *testing.T) {
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

	args := make([][]interface{}, 0)
	args = append(args, []interface{}{"set", testKey, testValue})
	args = append(args, []interface{}{"get", testKey})
	args = append(args, []interface{}{"exists", testKey})
	args = append(args, []interface{}{"del", testKey})
	args = append(args, []interface{}{"exists", testKey})

	cmds, err := proxy.Pipeline(args)
	if err != nil {
		t.Fatal(err)
	}
	for _, cmd := range cmds {
		t.Log(cmd.String(), cmd.Err())
	}
}
