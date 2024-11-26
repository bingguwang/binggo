package main

import (
	"context"
	"fmt"
	rv8 "github.com/go-redis/redis/v8"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	RedisCache struct {
		Host string
		Type string
		Pass string
	}
	BloomFilter struct {
		RedisKey  string
		BitSize   uint
		HashFuncs uint
	}
}

/*
*

	其实zero里是有布隆过滤器的，实现的是方式就是redis的bitmap
	所以zero的布隆过滤器是支持分布式的
*/
var (
	redisClient *redis.Redis
	config      Config
)

func main() {
	// 加载配置
	config = Config{
		RedisCache: struct {
			Host string
			Type string
			Pass string
		}{
			Host: "192.168.0.66:36379",
			Type: "node",
			Pass: "123456",
		},
		BloomFilter: struct {
			RedisKey  string
			BitSize   uint
			HashFuncs uint
		}{
			RedisKey:  "my-bloom-filter",
			BitSize:   10000,
			HashFuncs: 4,
		},
	}

	var err error
	// 初始化 Redis 客户端
	redisClient, err = redis.NewRedis(redis.RedisConf{
		Host: config.RedisCache.Host,
		Type: config.RedisCache.Type,
		Pass: config.RedisCache.Pass,
	})
	if err != nil {
		panic(err.Error())
	}

	// 初始化布隆过滤器
	bloomFilter := bloom.New(redisClient, config.BloomFilter.RedisKey, config.BloomFilter.BitSize)

	// 示例：添加和检查布隆过滤器
	key := "example-key"
	if err := initBloom(bloomFilter, key); err != nil {
		panic(err.Error())
	}

	// 示例：模拟查询和更新布隆过滤器
	simulateQueryAndUpdate(bloomFilter, key)
	go RebuildBloom()
	time.Sleep(5 * time.Second)
}

func initBloom(bloomFilter *bloom.Filter, key string) error {
	exists, err := bloomFilter.ExistsCtx(context.Background(), []byte(key))
	if err != nil {
		return err
	}
	if !exists {
		fmt.Println("Key", key, "不存在过滤器中，添加它到过滤器里")
		if err := bloomFilter.Add([]byte(key)); err != nil {
			log.Fatalf("Failed to add key to bloom filter: %v", err)
		}
	} else {
		fmt.Println("Key", key, "已存在过滤器中")
	}
	return nil
}

func simulateQueryAndUpdate(filter *bloom.Filter, key string) {
	ctx := context.Background()
	exists, err := filter.ExistsCtx(ctx, []byte(key))
	if err != nil {
		panic(err.Error())
	}

	if !exists {
		// 布隆过滤器认为该键一定不存在，我们直接返回不存在
		fmt.Println("Key not found in bloom filter, returning not found.")
		return
	}

	// 布隆过滤器认为该键可能存在，我们查询数据库进行确认
	fmt.Println("Possible key found in bloom filter, querying DB...")

	foundInDB := fakeDBQuery(key) // 模拟从数据库查询
	if !foundInDB {
		// 如果数据库中也不存在，则可以考虑是否需要将其添加到布隆过滤器，防止未来重复查询
		// 通常，我们不会在这种情况下更新布隆过滤器，因为这是假阳性的处理
		fmt.Println("Key not found in DB, but bloom filter thought it existed (false positive).")
	} else {
		// 如果数据库中存在，处理相应的业务逻辑
		fmt.Println("Key found in DB, processing business logic.")
	}
}

// 假设这个函数模拟从数据库查询，并返回键是否存在
func fakeDBQuery(key string) bool {
	// 模拟数据库查询逻辑，这里返回false表示数据库中不存在该键
	return false
}

/**
	由于数据库和缓存的情况会发生变化，所以会出现脏数据，但是我们又不能直接设置过滤器里的key为0
所以采用重建的方式来更新布隆过滤器
*/

func RebuildBloom() {
	rdb := rv8.NewClient(&rv8.Options{
		Addr:     config.RedisCache.Host, // Redis 服务器地址
		Password: config.RedisCache.Pass, // Redis 服务器密码
	})
	ctx := context.Background()

	// 定期重建布隆过滤器
	for {
		//time.Sleep(24 * time.Hour) // 每24小时重建一次
		time.Sleep(time.Second)

		// 1. 创建新的布隆过滤器
		bloomFilter := bloom.New(redisClient, "my-bloom-filter-new", config.BloomFilter.BitSize)
		if bloomFilter == nil {
			fmt.Println("Error creating new bloom filter ")
			continue
		}

		// 2. 从数据库中读取所有数据，并插入到新的布隆过滤器中
		// 这里假设有一个函数 getAllKeys 从数据库中获取所有键
		getAllKeys := func() []string {
			return []string{"1", "2", "3"}
		}
		allKeys := getAllKeys()
		for _, key := range allKeys {
			bloomFilter.AddCtx(ctx, []byte(key))
		}

		// 3. 原子替换旧的布隆过滤器

		// 执行 RENAME 命令
		err := rdb.Rename(ctx, "my-bloom-filter-new", "my-bloom-filter").Err()
		if err != nil {
			fmt.Println("Error renaming bloom filter:", err)
			continue
		}

		// 4. 删除旧的布隆过滤器
		err = rdb.Del(ctx, "my-bloom-filter-new").Err()
		if err != nil {
			fmt.Println("Error deleting old bloom filter:", err)
		}
	}
}
