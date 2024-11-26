package main

import (
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

func main() {
	var c kq.KqConf
	conf.MustLoad("config.yaml", &c)
	fmt.Println(c.Brokers)

	// 消费者组都在配置文件里配了，直接传入获的一个消费队列
	q := kq.MustNewQueue(c, kq.WithHandle(func(k, v string) error {

		fmt.Printf("=> %s\n", v)
		return nil
	}))
	defer q.Stop()
	q.Start()

}
