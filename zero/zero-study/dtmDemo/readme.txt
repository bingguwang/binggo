
dtm是一个分布式框架，处理分布式事务场景，保证数据的一致性
常见的分布式事务场景
订单系统：需要保证创建订单和扣减库存要么同时成功，要么同时回滚
跨行转账场景：数据不在一个数据库，但需要保证余额扣减和余额增加要么同时成功，要么同时失败
积分兑换场景：需要保证积分扣减和权益增加同时成功，或者同时失败
出行订票场景：需要在第三方系统同时定几张票，要么同时成功，要么全部取消
金融交易：确保资金转账等操作的一致性和可靠性。
跨服务协调：在微服务架构中，确保跨服务调用的一致性。

数据库与缓存一致性： dtm 的二阶段消息，能够保证数据库更新操作，和缓存更新/删除操作的原子性
秒杀系统： dtm 能够保证秒杀场景下，创建的订单量与库存扣减数量完全一样，无需后续的人工校准
多种存储组合： dtm 已支持数据库、Redis、Mongo 等多种存储，可以将它们组合为一个全局事务，保证数据的一致性


zero如何使用dtm:

先要搭建一下dtm并启动
git clone https://github.com/dtm-labs/dtmdriver-clients
config.yaml是dtm的配置文件，(需要模板配置文件去https://github.com/dtm-labs/dtm/blob/main/conf.sample.yml)
修改下zero相关的配置;
可以看下config.yaml文件。这个文件就是成功连接上了dtm的修改版

子事务屏障需要用到数据库，需要先创建个数据库 dtm_barrier，sql文件也给出了


运行dtm：
然后执行go run main.go -c指定配置文件config.yaml





也可以docker来搭建dtm

需要将此配置文件挂载到dtm docker容器里
然后写好docker-compose文件来运行dtm
但是，目前docker运行的dtm好像有问题，会出现错误 Error 1146 (42S02): Table 'dtm.kv' doesn't exist
