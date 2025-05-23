所谓的集群模式，就是多个nsqd节点，在发消息时，会同时发到集群里的所有nsqd节点对应的同名topic里，
也就是每个topic要对应多个生产者
从而达到高可用
var TopicProducers map[string][]*nsq.Producer
这是比较核心的结构，每个topic可以对应多个生产者

这里实现了一个集群案例，生产者和消费者在集群模式需要的操作

nsq的channel是什么时候创建的？
>在创建producer时，并不能设置topic对应的channel数，因为此时channel还不能创建出来
>当第一个消费者连接到 NSQD（消息队列节点）时，如果指定的 Topic 和 Channel 不存在，NSQD 便会自动创建它们
>因此，生产者在发送消息时不需要关心 Topic 和 Channel 是否存在

消费者一次可以处理的最大消息数量是 max-in-flight


集群模式保证了高可用，但是有个缺点，就是消费者端的重复消费问题,我执行例子也确实发现存在此问题
连接topic对应的所有 nsqd 的 consumer 将会收到多条重复消息，如何解决？？
* 消息去重
  在消费者端维护一个已处理消息的记录，比如使用缓存或数据库，每次处理消息前先检查这个记录，如果已经处理过则跳过。这种方法适用于消息幂等性较强的场景。
* 消息确认机制
  在 NSQ 中，消费者可以向 NSQ 提交消息确认（ACK），表示该消息已经被成功处理。NSQ 在收到 ACK 后会删除消息，确保不会再次传送给同一个消费者。如果消费者处理消息失败或者超时，则不提交 ACK，NSQ 会将消息重新发送给其他消费者。通过这种方式，可以确保消息只会被处理一次。
* 使用 NSQ 的保证交付（Guaranteed Delivery）模式
  NSQ 提供了 Guaranteed Delivery 模式，可以保证消息至少被消费一次。在这种模式下，NSQ 会在消息传送前存储消息，直到收到消费者的 ACK 才会删除消息。这种模式可以防止消息丢失，但会增加延迟和消耗存储空间。
* 消费者 ACK 时标记消息ID
  消费者在 ACK 消息时可以将消息的 ID 记录下来，如果后续收到相同 ID 的消息，就可以知道是重复消息。这需要在消费者端自行实现消息 ID 的记录和比对逻辑。
* 定期清理已处理消息的记录
  如果采用消息去重的方式，消费者可以定期清理已处理消息的记录，避免记录无限增长导致内存或存储资源耗尽。



