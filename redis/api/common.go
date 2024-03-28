package api

import (
	"context"
	"time"
)

/**
 * @Description: 设置一个键的超时时间
 * @receiver p
 * @param key: 键名
 * @param expiration: 超时时间
 * @return error
 */
func (p *RedisProxyer) Expire(key string, expiration time.Duration) error {
	ctx := context.Background()
	return p.Pool.Expire(ctx, key, expiration).Err()
}

/**
 * @Description: 设置一个键在指定时间超时
 * @receiver p
 * @param key: 键名
 * @param expiration: 超时时刻
 * @return error
 */
func (p *RedisProxyer) ExpireAt(key string, tm time.Time) error {
	ctx := context.Background()
	return p.Pool.ExpireAt(ctx, key, tm).Err()
}

/**
 * @Description: 用于迭代数据库中的键
 * @receiver p
 * @param cursor: 游标
 * @param match: 匹配模式，例如"somekey:*"
 * @param count: 返回元素个数
 * @return keys 查询到的键名
 * @return cur 游标新的位置，如果返回0表示迭代已结束
 * @return error 错误
 */
func (p *RedisProxyer) Scan(cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	ctx := context.Background()
	return p.Pool.Scan(ctx, cursor, match, count).Result()
}
