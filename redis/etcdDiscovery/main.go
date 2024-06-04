package main

import (
	register "binggo/etcd/registerDiscovery"
	"binggo/utils"
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

const (
	redisToRegister = "192.168.0.58"
)

func main() {
	exitChan := make(chan string)
	defer close(exitChan)
	ctx, cancel := context.WithCancel(context.Background())

	// 注册协程，一直开启直到不用时注销
	go serviceRegister(ctx, exitChan)
LOOP:
	for {
		select {
		case reason := <-exitChan: //出现错误则退出主线程
			fmt.Println("exit:", reason)
			cancel()
			break LOOP
		}
	}
	fmt.Println("sentinel daemon exit")

}

/** serviceRegister
* @Description: 向etcd注册redis服务
* @param ctx
* @param exitChan
 */
func serviceRegister(ctx context.Context, exitChan chan<- string) {
	defer wg.Done()

	fmt.Println("向etcd注册redis服务")
	uuid, err := utils.GetNodeUUID()
	if err != nil {
		fmt.Println("get node uuid failed")
		exitChan <- "exit"
		return
	}

	/*ip, err := utils.GetOutBoundIP()
	if err != nil {
		fmt.Println("get bound ip failed")
		exitChan <- "exit"
		return
	}*/
	ip := redisToRegister
	redisInfo := &RedisInfo{
		Name:     "mymaster",
		Password: "123456",
		CommonInfo: register.CommonInfo{
			Host: ip,
			//Port: "26381",
			Port: "26382",
		},
	}
	redisRegister, err := RegRedis(register.RedisServiceName, uuid, redisInfo, 5)
	if err != nil {
		fmt.Println("register failed", err.Error())
		exitChan <- "exit"
		return
	}
	defer redisRegister.UnRegister()

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("register ctx done")
			break LOOP
		}
	}

	fmt.Println("sentinel service end")
}
