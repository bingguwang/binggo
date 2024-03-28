package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

// const addr = "localhost:6379"
const addr = "192.168.2.130:6382"

func ExampleClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

}

// redis 默认是RDB持久化的，因为持久化所以不会出现内存崩溃的情况

func main() {
	ExampleClient()

	//count := 1000000000
	count := 50

	Write(count)

	chOs := make(chan os.Signal, 1)
	signal.Notify(chOs, syscall.SIGTERM, syscall.SIGINT)

	rand.Seed(time.Now().UnixNano())
	for {
		select {
		case <-time.After(3 * time.Second):
			randomInt := rand.Intn(count)
			k := "key" + strconv.Itoa(randomInt)
			fmt.Println(k)
			val2, err := rdb.Get(ctx, k).Result()
			if err == redis.Nil {
				fmt.Println("key does not exist")
			} else if err != nil {
				panic(err)
			} else {
				fmt.Println("key", val2)
			}
		case <-chOs:
			return
		}
	}
}

func Write(count int) {
	for i := 0; i < count; i++ {
		// set a value with a cost of 1
		val, _ := json.Marshal(&IpcInfo{})
		val2, _ := json.Marshal(&OrganizationInfo{})

		if err := rdb.Set(ctx, "key"+strconv.Itoa(i), val, 0).Err(); err != nil {
			panic(err)
		}
		if err := rdb.Set(ctx, "key2"+strconv.Itoa(i), val2, 0).Err(); err != nil {
			panic(err)
		}
	}
}

type IpcInfo struct {
	Id                 string `json:"id"`
	Uuid               string `json:"uuid"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	Ip                 string `json:"ip"`
	Vendor             string `json:"vendor"`
	Model              string `json:"model"`
	PlatformId         string `json:"platform_id"`
	Status             string `json:"status"`
	RecordCompleteRate string `json:"record_complete_rate"`
	DomainId           string `json:"domain_id"`
	Appearance         string `json:"appearance"`
}

type OrganizationInfo struct {
	Id       string `json:"id"`
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Parent   string `json:"parent"`
	Platform string `json:"platform"`
	Path     string `json:"path"`
	DomainId string `json:"domain_id"`
}
