## 如何注册
首先，先得定义一个结构用于注册，这个结构需要有注册相关的必须属性
一个注册得有这些东西：

> 服务名
> 服务信息
> 注册时间
> 注册超时时间
> 操作etcd的client
> 唯一标识注册的uuid
> 租约id


然后注册其实是这样一个过程》》

因为etcd其实是个分布式的kv存储中间件
所以需要想好服务注册时候的key是啥？

> 一个比较好的key，例如 `**服务前缀/服务名称/uuid/**`


有了key之后，就是value存什么？value当然是节点相关的信息了，也就是node ip
存的过程大致如下：

* 创建一个租约
* put服务的key-value并设置租约
* 调用KeepAlive给租约续约，返回的租约回复通道应该存到注册的结构里

于是一个注册结构里应该有这些属性：

> 服务名
> 服务信息
> 注册时间
> 注册超时时间
> 操作etcd的client
> 唯一标识注册的uuid
> 租约id
> 租约回复通道

由于在，注册过程中可能会发生服务的一些信息变换，所以在注册的过程里可能还需要对这些信息进行更新,
比如之前定义在注册结构里的服务信息，这是可能会发生变换的


## 如何注销
同样的，注销就是撤销租约即可，这样，etcd里和租约绑定的key会被立刻删除



## 如何发现
要发现，同样先定义一个结构体，这个结构体应有如下属性：
> 服务名
> 服务信息
> 操作etcd的client
> 操作的锁

其中服务信息应该要可以存多个服务的服务信息，这点和注册不一样，因为发现这应该是可以看到多个服务的信息的

发现到服务之后，还应该持续关注这些服务的信息是否发生变更，这一点可以使用WatchWithPrefix实现
发现变更及时更新到发现结构的服务信息里
