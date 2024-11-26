package bloom

import (
	"fmt"
	"github.com/willf/bloom"
	"testing"
)

// 模拟缓存和数据库
var cache = make(map[string]string)
var database = map[string]string{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
}

// 初始化布隆过滤器
var bf = bloom.New(1000, 4)

func init() {
	// 将现有的数据库键插入布隆过滤器，在插入新数据的时候得更新布隆过滤器
	for key := range database {
		bf.Add([]byte(key))
	}
}

func query(key string) (string, error) {
	// 先查询布隆过滤器
	if !bf.Test([]byte(key)) {
		return "", fmt.Errorf("key not found")
	}

	// 查询缓存
	if value, found := cache[key]; found {
		return value, nil
	}

	// 查询数据库
	if value, found := database[key]; found {
		// 更新缓存
		cache[key] = value
		return value, nil
	}

	// 如果数据库中也没有，返回未找到
	return "", fmt.Errorf("key not found")
}

func TestBloom(t *testing.T) {
	// 示例查询
	keys := []string{"key1", "key2", "key4"}

	for _, key := range keys {
		value, err := query(key)
		if err != nil {
			fmt.Printf("Query for %s failed: %v\n", key, err)
		} else {
			fmt.Printf("Query for %s succeeded: %s\n", key, value)
		}
	}
}
