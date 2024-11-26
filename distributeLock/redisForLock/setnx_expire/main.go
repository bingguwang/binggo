package main

import (
	"binggo/distributeLock/redisForLock"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var (
	ctx = context.Background()
)

// 利用setnx， 如果 key不存在，则SETNX成功返回1，如果这个key已经存在了，则返回0
// 先用setnx来抢锁，如果抢到之后，再用expire给锁设置一个过期时间，防止锁忘记了释放。

/*
业务执行的时间需要把控好，不然过期时间不好确定
而且还要保证分布式下各个节点的时间是同步的
*/
func main() {
	client := redisForLock.GetClient()
	keyResourceID := "my_resource_lock"
	lockValue := uuid.New().String()

	//set key value px milliseconds nx
	boolCmd := client.SetNX(ctx, keyResourceID, lockValue, 10*time.Second)
	if boolCmd.Val() { // 拿到锁了
		done := make(chan struct{}) // 用于业务协程和续约协程之间通信

		// 锁续约机制, 增加一个协程在业务完成之前都去更新一下过期时间
		go func() {
			ticker := time.NewTicker(8 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					// todo 也不是原子操作，所以还是需要用lua脚本才合理
					if client.Get(ctx, keyResourceID).Val() == lockValue {
						client.Expire(ctx, keyResourceID, 10*time.Second)
					}
					/*
						// Lua 脚本检查锁的值并续约
						luaScript := `
							if redis.call("GET", KEYS[1]) == ARGV[1] then
								return redis.call("EXPIRE", KEYS[1], ARGV[2])
							else
								return 0
							end
						`
						expireResult, err := client.Eval(ctx, luaScript, []string{keyResourceID}, lockValue, 10).Result()
						if err != nil {
							fmt.Println("Error extending lock:", err)
						} else if expireResult == int64(0) {
							fmt.Println("Failed to extend lock: lock value mismatch")
						}
					*/
				case <-done:
					return
				}
			}
		}()

		// 处理业务
		go func() {
			defer func() {
				// 释放锁时检查锁的值是否匹配
				// 检查当前锁的值是否与 lockValue 匹配，以确保只释放自己获取的锁。
				/*
					因为如果客户端A在业务处理过程中由于某种原因挂起或延迟，而过期时间到了，锁会自动释放。
					这时，客户端B可能会获取到这个锁。
					如果客户端A在业务处理完成后尝试释放锁，它可能会错误地删除客户端B的锁，导致客户端B的操作被中断。
				*/
				if client.Get(ctx, keyResourceID).Val() == lockValue {
					// todo 这里Get到Del之间不是原子操作，会有竞争条件，比如clientA走到这里的时候过期了，clientB此时获取到了锁，
					// todo 然后clientA执行Del,删除了clientB获取的锁
					// todo 可能会有点疑惑，为啥这里的还会出现clientA的锁过期的情况？我的clientA不是有续约机制吗？
					// todo 事实上，续约机制不是完全可靠的，可能网络少延迟一下就预约失败了
					// todo 所以对于这里不是原子性的操作，最好还是使用lua脚本来保证原子性
					client.Del(ctx, keyResourceID)
				}

				/*// Lua 脚本确保检查和删除操作的原子性
				luaScript := `
					if redis.call("GET", KEYS[1]) == ARGV[1] then
						return redis.call("DEL", KEYS[1])
					else
						return 0
					end
				`
				_, err := client.Eval(ctx, luaScript, []string{keyResourceID}, lockValue).Result()
				if err != nil {
					fmt.Println("Error releasing lock:", err)
				}*/

				close(done) // 业务结束通知续约协程退出
			}()

			time.Sleep(12 * time.Second)
			fmt.Println("处理业务")
		}()
	} else {
		fmt.Println("锁被其他人占用")
	}

	// 实时打印锁key的剩余时间
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				ttl := client.TTL(ctx, keyResourceID).Val()
				if ttl.Seconds() > 0 {
					fmt.Printf("Remaining TTL: %.0f seconds\n", ttl.Seconds())
				} else {
					fmt.Println("Key has expired or does not exist")
					return
				}
			}
		}
	}()

	// 为了防止程序结束时 goroutine 被杀死
	time.Sleep(1000 * time.Second)
	client.Close()
}
