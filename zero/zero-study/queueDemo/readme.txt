消息队列有两种：

DQ（Dead Queue）：死信队列，用于存储和处理无法正常消费的消息，帮助排查和解决问题。
 依赖于 beanstalkd，分布式，可存储，延迟、定时设置，关机重启可以重新执行，消息不会丢失，
 使用非常简单，go-queue中使用了redis setnx保证了每条消息只被消费一次，使用场景主要是用来做日常任务使用

KQ（Kafka Queue）：基于Kafka实现的消息队列，适用于高吞吐量、持久化、分布式和实时数据


简单了解下 beanstalkd ，go-zero的dq是基于此实现的
https://segmentfault.com/a/1190000042205622

job：任务单元；
tube：任务队列，存储统一类型 job。producer 和 consumer 操作对象；
producer：job 生产者，通过 put 将 job 加入一个 tube；
consumer：job 消费者，通过 reserve/release/bury/delete 来获取 job 或改变 job 的状态



由于kafka本身不支持延时消费， 而beanstalkd支持定时和延时操作。
所以基于实现了dq，另外还加入 redis 保证消费唯一性。


于是当有延时、定时任务执行的场景时，使用dq
当有 异步、批量任务执行的场景时，使用kq






