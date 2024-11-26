package tcp

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net"
	"regexp"
	"sync"
	"testing"
	"time"
)

// 测试TCP端口是否开启
func testTCPPort(host string, port int, timeout time.Duration) bool {
	// 构造TCP地址
	address := fmt.Sprintf("%s:%d", host, port)

	// 创建一个TCP连接
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		// 如果连接失败，则认为端口未开启
		fmt.Printf("连接失败: %v\n", err)
		return false
	}

	// 关闭连接
	conn.Close()

	// 如果连接成功，则认为端口开启
	return true
}

func TestPort(t *testing.T) {
	host := "192.168.2.142"    // 需要测试的服务器地址
	port := 2020               // 需要测试的端口号
	timeout := 5 * time.Second // 设置超时时间

	if testTCPPort(host, port, timeout) {
		fmt.Printf("端口 %d 在 %s 上是开启的\n", port, host)
	} else {
		fmt.Printf("端口 %d 在 %s 上未开启或无法连接\n", port, host)
	}
}

func TestCase1(t *testing.T) {
	res, err := CheckPort("192.168.2.133", 7777, 3*time.Second)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res)
}

func CheckPort(host string, port int, timeout time.Duration) (bool, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}
func TestRg(t *testing.T) {
	vcsRedisClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.154:6379",
		Password: "",
		DB:       1,
	})
	ctx := context.Background()
	OrganizationUuid := "fbcc305c-fde3-4a3f-ba74-bc85f200c641"
	result, err := vcsRedisClient.HGet(ctx, "orgList", OrganizationUuid).Result()
	if err != nil {
		fmt.Sprintf("vcs组织[%s]不存在", OrganizationUuid)
		return
	}
	fmt.Println(result)
	/*if err := vcsRedisClient.HSet(ctx, "IPC", "XXX", "aaaa").Err(); err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	if err := vcsRedisClient.HDel(ctx, "IPC", "XXX").Err(); err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println("不存在的修改，不会爆错")*/
}
func TestName(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	timer := time.NewTimer(3 * time.Second)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(50 * time.Millisecond)
			fmt.Println("打印")
			select {
			case <-timer.C:
				fmt.Println("退出监听")
				return
			default:
			}
		}
	}()
	wg.Wait()
	fmt.Println("退出")
}

func TestNamedd(t *testing.T) {

	input := "Message<vcsChannel:ipcList|1b627762-2d34-42f1-8cf5-74532694а96b|ADD>"

	// 使用正则表达式提取
	re := regexp.MustCompile(`vcsChannel:([^>]+)>`)
	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		fmt.Println(match[0])
		result := match[1]
		fmt.Println(result)
		fmt.Printf("%c", result[0])
	} else {
		fmt.Println("No match found")
	}
}
