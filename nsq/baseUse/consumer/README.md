
```go
type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// 消费者要怎么处理消息就在这里
	fmt.Printf("从%s消费消息；%s\n", m.NSQDAddress, time.Now().String())

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return nil
}

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic", "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	// 为消费者收到的消息设置handler
	consumer.AddHandler(&myMessageHandler{})
	//consumer.AddConcurrentHandlers 可以同时开启多个协程执行handler处理消息

	// 通过 nsqlookupd 来发现nsq实例
	err = consumer.ConnectToNSQLookupd("192.168.2.130:41610") // 传入的地址是nsqlookupd的地址
	if err != nil {
		log.Fatal(err)
	}

	// 执行的时候发现每60s会去 querying nsqlookupd http://192.168.2.130:41610/lookup?topic=topic名称
	// 就是查询一次所有含有此topic的nsqd节点
	/**
	16:01:09
	16:02:16
	16:03:16
	16:04:16
	*/

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}

```

调用`ConnectToNSQLookupd`时，执行的时候发现每60s会执行一次

`querying nsqlookupd http://192.168.2.130:41610/lookup?topic=topic名称`
此操作就是查询一次所有含有此topic的nsqd节点


而这步在源码里也有体现


```go
func (r *Consumer) lookupdLoop() {
	// add some jitter so that multiple consumers discovering the same topic,
	// when restarted at the same time, dont all connect at once.
	r.rngMtx.Lock()
	jitter := time.Duration(int64(r.rng.Float64() *
		r.config.LookupdPollJitter * float64(r.config.LookupdPollInterval)))
	r.rngMtx.Unlock()
	var ticker *time.Ticker

	select {
	case <-time.After(jitter):
	case <-r.exitChan:
		goto exit
	}

	ticker = time.NewTicker(r.config.LookupdPollInterval)

	for {
		select {
		case <-ticker.C:
			r.queryLookupd()
		case <-r.lookupdRecheckChan:
			r.queryLookupd()
		case <-r.exitChan:
			goto exit
		}
	}

exit:
	if ticker != nil {
		ticker.Stop()
	}
	r.log(LogLevelInfo, "exiting lookupdLoop")
	r.wg.Done()
}
```

这个方法是每个一定时间查询一次含有此topic的nsqd节点

时间间隔如下：

![image](assets/image-20240423161034-7r2ofhy.png)




消费者的一些参数
> - -max-in-flight ：一次可以处理的最大消息数量 
> - -msg-timeout ：单条消息的超时时间，默认一分钟，即消息投递后一分钟内未收到响应，则 nsq 会将这条消息 requeue 处理(再次投递)。
> - -max-msg-timeout ：nsqd 全局设置的最大超时时间，默认 15 分钟。
