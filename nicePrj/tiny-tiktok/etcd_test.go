package tiny_tiktok

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestEtcd(t *testing.T) {

	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	// 创建配置对象，指定server地址并设置超时时间
	// 这里因为我用的是windows系统 docker安装在虚拟中
	// 所以地址填的是虚拟机ip
	config = clientv3.Config{
		Endpoints:   []string{"http://192.168.2.130:2379"},
		Username:    "root",
		Password:    "123456",
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		// 只是测试一下，有错误就直接panic吧
		panic(err.Error())
	}
	fmt.Println("suc", client == nil)
	fmt.Println("suc", client.Password)

	// 读取/cron/jobs/为前缀的所有key
	if getResp, err := client.Get(context.TODO(), "/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else { // 获取成功, 我们遍历所有的kvs
		fmt.Println(getResp.Kvs)
	}
}
