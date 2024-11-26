package main

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"time"
)

/*
*
setnx_expire里的单机的redis实现的分布式锁存在单点故障，
即使Redis通过sentinel保证高可用，如果这个master节点由于某些原因发生了主从切换，那么还是会出现锁丢失的情况：
因为sentinel是主从复制的，如果master的key没有及时同步到slave，而master就宕机了，那slave升级为master时候，这个key也是没有的，锁丢失了

所以最可靠的话还得的是redis集群实现的分布式锁

而基于集群的分布式锁用到了redlock算法:
此算法保证了 任何时刻，只有一个客户端持有锁，且不会产生死锁，只要大多数redis节点能够正常工作，客户端端都能获取和释放锁。

RedLock：
所有的节点之间不搞注册复制那套
客户端怎样算获取到锁了？

	需要满足从大多数(超过半数)的节点上都获取到锁了，才是真的获取到锁了，
	否则就是没有成功获取锁，需要在所有的节点(包括并没有成功获取锁的节点)上进行释放锁的操作
*/
var ctx = context.Background()

/*
func main() {
	// Connect to redis.
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	defer client.Close()

	// 配置 Redis 客户端
	client1 := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	client2 := redis.NewClient(&redis.Options{Addr: "localhost:6380"})
	client3 := redis.NewClient(&redis.Options{Addr: "localhost:6381"})

	// 创建 Redlock 实例
	locker1 := redislock.New(client1)
	locker2 := redislock.New(client2)
	locker3 := redislock.New(client3)

	// 获取锁
	key := "my_resource_lock"
	value := uuid.New().String()
	lock, err := locker1.Obtain(ctx, key, 10*time.Second, nil)
	if err == redislock.ErrNotObtained {
		fmt.Println("锁被其他人占用")
		return
	} else if err != nil {
		fmt.Println("获取锁失败:", err)
		return
	}

	fmt.Println("获取到锁")

	// 锁续约机制
	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(8 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := lock.Refresh(ctx, 10*time.Second, nil)
				if err != nil {
					fmt.Println("锁续约失败:", err)
					return
				}
			case <-done:
				return
			}
		}
	}()

	// 处理业务
	go func() {
		defer func() {
			if err := lock.Release(ctx); err != nil {
				fmt.Println("释放锁失败:", err)
			}
			close(done)
		}()

		time.Sleep(3 * time.Second)
		fmt.Println("处理业务")
	}()

	// 为了防止程序结束时 goroutine 被杀死
	time.Sleep(10 * time.Second)
	client1.Close()
	client2.Close()
	client3.Close()
}
*/

func main() {
	// 创建 Redis 客户端
	client1 := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.1:5378",
		Password: "a123456",
		DB:       0,
	})

	client2 := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.1:5379",
		Password: "a123456",
		DB:       0,
	})

	client3 := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.1:5380",
		Password: "a123456",
		DB:       0,
	})

	// 创建 RedisLock 实例
	locker1 := redislock.New(client1)
	locker2 := redislock.New(client2)
	locker3 := redislock.New(client3)

	// 尝试获取分布式锁
	resourceName := "REDLOCK_KEY"
	ctx := context.Background()
	var lock1, lock2, lock3 *redislock.Lock
	var err error

	for i := 0; i < 3; i++ {
		lock1, lock2, lock3, err = tryAcquireLock(ctx, locker1, locker2, locker3, resourceName, 10*time.Second)
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	if err != nil {
		fmt.Println("Failed to acquire lock:", err)
		return
	}

	fmt.Println("Lock acquired, doing something...")

	// 处理业务

	// 释放锁
	if lock1 != nil {
		if err := lock1.Release(ctx); err != nil {
			fmt.Println("Failed to release lock1:", err)
		}
	}
	if lock2 != nil {
		if err := lock2.Release(ctx); err != nil {
			fmt.Println("Failed to release lock2:", err)
		}
	}
	if lock3 != nil {
		if err := lock3.Release(ctx); err != nil {
			fmt.Println("Failed to release lock3:", err)
		}
	}

	fmt.Println("Lock released")
}

// tryAcquireLock 尝试获取分布式锁
func tryAcquireLock(ctx context.Context, locker1, locker2, locker3 *redislock.Client, key string, ttl time.Duration) (*redislock.Lock, *redislock.Lock, *redislock.Lock, error) {
	lock1, err1 := locker1.Obtain(ctx, key, ttl, nil)
	lock2, err2 := locker2.Obtain(ctx, key, ttl, nil)
	lock3, err3 := locker3.Obtain(ctx, key, ttl, nil)

	if err1 == nil && err2 == nil && err3 == nil {
		return lock1, lock2, lock3, nil
	}

	if lock1 != nil {
		lock1.Release(ctx)
	}
	if lock2 != nil {
		lock2.Release(ctx)
	}
	if lock3 != nil {
		lock3.Release(ctx)

	}

	if err1 != nil {
		return nil, nil, nil, err1
	}
	if err2 != nil {
		return nil, nil, nil, err2
	}
	return nil, nil, nil, err3
}
