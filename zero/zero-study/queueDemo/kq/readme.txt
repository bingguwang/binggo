
zero的kq是使用kafka实现的




消息队列的一般作用就是解耦，异步处理，削峰

kafka的消费方式是由consumer去pull的模式

每个消息都应该知道自己要发到哪个topic里
每个topic都有一个默认的partition，每个topic可以设置多个partition，这些partition可以在不同机器上（kafka集群里机器也称为broker）
partition在磁盘上其实就是一个目录

消费的时候一般是配合消费组来消费，每个组里可以有多个消费者，但是组内每条消息只有一个消费者可以消费

举例子：每来一个消费者组就会分一个榴莲(一个topic)，这个榴莲是copy的，保证每个组分到的榴莲一模一样)
，一个榴莲分开成一片一片，每一片都不一样，然后组里的每个消费者分一片，每个消费者只能吃自己组分到的一片或几片榴莲，
是的，组内饭量大的可以吃好几片，但每个组里的每片榴莲只能给一个人吃（不可能某时间段里一片榴莲给多个人吃）
每片榴莲就是一个partition，一个组里的所有榴莲片加起来是一个完整的榴莲(也就是一个topic)
每片榴莲是不一样的(每个partition里的数据是不一样的)


那问题来了：现在假设榴莲片增加了咋办？谁去吃？某个消费者down了，或者加了个消费者，组里的每片榴莲又该如何分配？
答：重新分配，这时候你没吃完的那片榴莲可能分给别人吃了，然后你吃别人没吃完的榴莲片


那副本和多节点是啥？
副本指的是分区的副本，一个分区有多副本，副本又分为一个leader和其余的follower
在写入消息时，在消费消息时，都是找的leader去操作的！副本只是同步leader的数据而已，在leader挂后，follower才顶上作为新的leader
每个分区的副本都分布在不同的机器节点broker上
因为partition分布在不同的机器上，所以可以通过扩展机器来增大并发能力


来看下partition啥样吧：
每个partition是一个目录，下面有一个个的segment目录，在这之下有index文件、.log文件、.timeindex文件
.log文件就是存消息的文件，分区消费到哪里了有其他两个文件检索
对于消费过的消息不会马上删除(会定期去删除)


在写入数据时，每个topic有多个分区，咋知道写到哪个partition里呢？
首先这是可以在写入时指定分区的，不指定分区如果设置了数据的key，就根据key hash出一个分区来写入
如果没指定分区和key，就是轮训的方式选一个partition写入


https://www.zhihu.com/search?type=content&q=kafka


实践一下：
按照所写的docker-compose创建的一个 broker是2，副本数是2，分区数是3的kafka集群：
创建了一个名为kq的topic
2.130上有2个kafka
0.66有1个kafka
然后看下topic目录
在2.130上，/data/kafka1/ 下没有创建topic目录
/data/kafka2/ 下如下：
/data/kafka2/
└── kafka-logs-246d454d6023
    ├── cleaner-offset-checkpoint
    ├── kq-0
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── kq-1
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── kq-2
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── log-start-offset-checkpoint
    ├── meta.properties
    ├── recovery-point-offset-checkpoint
    └── replication-offset-checkpoint

三个kq目录就是分区，这三个分区里的数据是不一样的

在0.66上，/data/kafka1/如下：
/data/kafka1/
└── kafka-logs-24bc25c8d7c3
    ├── cleaner-offset-checkpoint
    ├── __consumer_offsets-0
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-1
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-10
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-11
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-12
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-13
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-14
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-15
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-16
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-17
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-18
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-19
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-2
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-20
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-21
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-22
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-23
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-24
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-25
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-26
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-27
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-28
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-29
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-3
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-30
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-31
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-32
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-33
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-34
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-35
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-36
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-37
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-38
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── __consumer_offsets-39
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── kq-0
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── kq-1
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── kq-2
    │       ├── 00000000000000000000.index
    │       ├── 00000000000000000000.log
    │       ├── 00000000000000000000.timeindex
    │       └── leader-epoch-checkpoint
    ├── log-start-offset-checkpoint
    ├── meta.properties
    ├── recovery-point-offset-checkpoint
    └── replication-offset-checkpoint

也可以看到三个kq目录。三个目录下的数据也是不一样的
因为副本数是2，所以2.130上的/data/kafka1/下没有topic目录
另外在0.66的/data/kafka1下看到了consumer相关的消息，在2.130上面没有，说明数据被写入到了一个分区里而已，
而且消费的时候只会取消费此分区的leader副本

follower如何知道leader消费到哪里了？follower会同步leader的相关文件，比如这里看到的各种checkpoint和index


go-queue怎么在消费和生产时指定分区？
目前看go-queue好像没有的提供相关接口，负载均衡啥的都封装死了

