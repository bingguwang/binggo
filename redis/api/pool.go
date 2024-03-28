package api

import (
	"context"
	"errors"
	"log"

	"github.com/go-redis/redis/v8"
)

/**
 * @Description: 分配redis存储池
 * @param masterName：master名称
 * @param SentinelAddrs：哨兵地址列表
 * @param MinIdleConns：最小连接数
 * @return *redis.ClusterClient：池操作对象
 * @return error
 */
func NewRedisPool(masterName, passwd string, sentinelAddrs []string, minIdleConns int) (*redis.ClusterClient, error) {
	option := &redis.FailoverOptions{
		MasterName:         masterName,
		SentinelAddrs:      sentinelAddrs,
		RouteByLatency:     false,        // 只读请求路由到最近的节点
		RouteRandomly:      true,         // 只读请求路由随机
		DB:                 0,            // redis数据库index
		MaxRetries:         0,            // 命令执行失败时，最多重试多少次，默认3次，-1 (not 0) 表示取消重试
		MinRetryBackoff:    0,            // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff:    0,            // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
		DialTimeout:        0,            // 连接建立超时时间，默认5秒，-1表示取消读超时
		ReadTimeout:        0,            // 读超时，默认3秒，-1表示取消读超时
		WriteTimeout:       0,            // 写超时，默认等于读超时
		PoolSize:           0,            // 连接池最大socket连接数，默认 runtime.GOMAXPROCS * 10
		MinIdleConns:       minIdleConns, // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量
		MaxConnAge:         0,            // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
		PoolTimeout:        0,            // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为ReadTimeout+1秒
		IdleTimeout:        0,            // 闲置超时，默认5分钟，-1表示取消闲置超时检查
		IdleCheckFrequency: 0,            // 空闲连接收割者进行空闲检查的频率，默认1分钟；-1禁用空闲连接收割者，但如果设置了IdleTimeout，客户端仍会丢弃空闲连接
	}
	if passwd != "" {
		option.Password = passwd
	}
	cli := redis.NewFailoverClusterClient(option)
	if cli == nil {
		return nil, errors.New("create client failed")
	}

	ctx := context.Background()
	err := cli.Ping(ctx).Err()
	if err != nil {
		log.Println("[cache] Ping:", err)
		return nil, err
	}

	return cli, nil
}
