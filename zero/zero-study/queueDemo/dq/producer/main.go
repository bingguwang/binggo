package main

import (
	"fmt"
	"github.com/zeromicro/go-queue/dq"
	"strconv"
	"time"
)

/*
*

job：任务单元；
tube：任务队列，存储统一类型 job。producer 和 consumer 操作对象；
producer：job 生产者，通过 put 将 job 加入一个 tube；
consumer：job 消费者，通过 reserve/release/bury/delete 来获取 job 或改变 job 的状态

生产者需要关注的是 往哪个tube加入任务，tube的任务类型是啥
*/
var ip = "192.168.2.130"

func main() {
	// [Producer] 多节点模式至少要2个节点
	producer := dq.NewProducer([]dq.Beanstalk{
		{
			Endpoint: ip + ":9093",
			Tube:     "tube1",
		},
		{
			Endpoint: ip + ":9094",
			Tube:     "tube2",
		},
	})

	for i := 1000; i < 1005; i++ {
		// 延迟到5秒后生产消息
		fmt.Println("延迟到5秒后生产消息")
		// 消息放入[]byte内
		_, err := producer.Delay([]byte(strconv.Itoa(i)), time.Second*15)
		if err != nil {
			fmt.Println(err)
		}
	}
	/**
	测试与总结:
	两个节点都正常开启，消费者端只消费一个节点，发现消费还是可以消费到完整的消息（但由于redis setnx没有出现重复消费），
	可见两个节点里都有一份完整的消息
	一旦数组里的节点有一个down了，消息生成都不会成功
	*/
}

// 单节点生产者
//func main() {
//	producer := dq.NewProducerNode("localhost:11300", "tube")
//
//	for i := 1000; i < 1005; i++ {
//		_, err := producer.Delay([]byte(strconv.Itoa(i)), time.Second*5)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//}
