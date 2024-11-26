package bloom

import (
	"context"
)

var ctx = context.Background()

//func TestC1(t *testing.T) {
//	rdb := redis.NewClient(&redis.Options{
//		Addr: "localhost:6379",
//	})
//
//	bf := redisbloom.NewClient(rdb)
//
//	// 初始化原有的布隆过滤器
//	_, err := bf.Create(ctx, "bloom_filter", 1000, 0.01)
//	if err != nil {
//		fmt.Println("Error creating bloom filter:", err)
//		return
//	}
//
//	// 模拟插入一些数据
//	bf.Add(ctx, "bloom_filter", "key1")
//	bf.Add(ctx, "bloom_filter", "key2")
//
//	// 定期重建布隆过滤器
//	for {
//		time.Sleep(24 * time.Hour) // 每24小时重建一次
//
//		// 1. 创建新的布隆过滤器
//		_, err := bf.Create(ctx, "bloom_filter_new", 1000, 0.01)
//		if err != nil {
//			fmt.Println("Error creating new bloom filter:", err)
//			continue
//		}
//
//		// 2. 从数据库中读取所有数据，并插入到新的布隆过滤器中
//		// 这里假设有一个函数 getAllKeys 从数据库中获取所有键
//		allKeys := getAllKeys()
//		for _, key := range allKeys {
//			bf.Add(ctx, "bloom_filter_new", key)
//		}
//
//		// 3. 原子替换旧的布隆过滤器
//		err = rdb.Rename(ctx, "bloom_filter_new", "bloom_filter").Err()
//		if err != nil {
//			fmt.Println("Error renaming bloom filter:", err)
//			continue
//		}
//
//		// 4. 删除旧的布隆过滤器
//		err = rdb.Del(ctx, "bloom_filter_old").Err()
//		if err != nil {
//			fmt.Println("Error deleting old bloom filter:", err)
//		}
//	}
//}
//
//// 模拟从数据库中获取所有键的函数
//func getAllKeys() []string {
//	return []string{"key1", "key2", "key3"}
//}
