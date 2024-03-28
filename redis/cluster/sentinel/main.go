package main

import (
	rd "binggo/redis"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

/**
在sentinel模式下，如何操作redis
因为只能在主节点写入，所有连接redis就必须是动态的
*/

const (
	node1         = "192.168.2.130:6381"
	node2         = "192.168.2.130:6382"
	node3         = "192.168.2.130:6383"
	sentinelNode1 = "192.168.2.130:26381"
	sentinelNode2 = "192.168.2.130:26382"
	sentinelNode3 = "192.168.2.130:26383"
	MasterName    = "mymaster"
)

var ctx = context.Background()

func main() {

	client := redis.NewFailoverClient(&redis.FailoverOptions{
		//MasterName: "myredis-1",
		MasterName:    MasterName,
		Password:      "123456",
		SentinelAddrs: []string{sentinelNode1, sentinelNode2, sentinelNode3},
		//SentinelAddrs: []string{node1, node2, node3},
	})

	/** 路由只读命令到从节点
	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
	   MasterName:    "mymaster",
	   SentinelAddrs: []string{node1, node2, node3},

	   // 你可以选择把只读命令路由到最近的节点，或者随机节点，二选一
	   // RouteByLatency: true,
	   // RouteRandomly: true,
	})
	*/
	if client == nil {
		panic("get client failed ")
	}
	fmt.Println("初始化成功")
	if setCmd := client.Set(ctx, "k1", "v1", 0); setCmd.Err() != nil {
		panic(setCmd.Err().Error())
	}
	fmt.Println("设置值成功")

	sentinel := redis.NewSentinelClient(&redis.Options{
		Addr: sentinelNode3,
	})

	addr, err := sentinel.GetMasterAddrByName(ctx, MasterName).Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("主节点:", addr)
	sentinels := sentinel.Sentinels(ctx, MasterName)
	res, _ := sentinels.Result()
	for _, r := range res {
		fmt.Println(rd.TojsonStr(r))
	}
}
