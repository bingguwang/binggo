package api

import (
	"context"
	"time"
)

// 获取队列长度
func (p *RedisProxyer) LLen(key string) (int64, error) {
	ctx := context.Background()
	return p.Pool.LLen(ctx, key).Result()
}

/**
 * @Description:获取列表指定范围内的元素
 * @receiver p
 * @param key：list键值
 * @param start：开始的索引，从0开始
 * @param stop：结束索引
 * @return []string：列表元素列表
 * @return error
 */
func (p *RedisProxyer) LRange(key string, start, stop int64) ([]string, error) {
	ctx := context.Background()
	return p.Pool.LRange(ctx, key, start, stop).Result()
}

/**
 * @Description:将一个或多个值插入到列表头部
 * @receiver p
 * @param key: list键值
 * @param values：值，按顺序插入；举例 lpush list 1 2 3, list元素为[3,2,1]
 * @return error
 */
func (p *RedisProxyer) LPush(key string, values ...interface{}) error {
	ctx := context.Background()
	return p.Pool.LPush(ctx, key, values...).Err()
}

// 移出并获取列表的第一个元素
func (p *RedisProxyer) LPop(key string) (string, error) {
	ctx := context.Background()
	return p.Pool.LPop(ctx, key).Result()
}

func (p *RedisProxyer) LPopInt(key string) (int, error) {
	ctx := context.Background()
	return p.Pool.LPop(ctx, key).Int()
}

/**
 * @Description: 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
 * @receiver p
 * @param timeout: 等待超时时间，最小为1秒
 * @param key：选中的list
 * @return []string：list的弹出元素, 格式为[list_key value]
 * @return error
 */
func (p *RedisProxyer) BLPop(timeout time.Duration, key string) ([]string, error) {
	ctx := context.Background()
	return p.Pool.BLPop(ctx, timeout, key).Result()
}

/**
 * @Description:将一个或多个值插入到列表尾部
 * @receiver p
 * @param key: list键值
 * @param values：值，按顺序插入；举例 rpush list 1 2 3, list元素为[1,2,3]
 * @return error
 */
func (p *RedisProxyer) RPush(key string, values ...interface{}) error {
	ctx := context.Background()
	return p.Pool.RPush(ctx, key, values...).Err()
}

// 移除列表的最后一个元素
func (p *RedisProxyer) RPop(key string) (string, error) {
	ctx := context.Background()
	return p.Pool.RPop(ctx, key).Result()
}

func (p *RedisProxyer) RPopInt(key string) (int, error) {
	ctx := context.Background()
	return p.Pool.RPop(ctx, key).Int()
}

/**
 * @Description: 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
 * @receiver p
 * @param timeout: 等待超时时间，最小为1秒
 * @param key：选中的list
 * @return []string：list的弹出元素, 格式为[list_key value]
 * @return error
 */
func (p *RedisProxyer) BRPop(timeout time.Duration, key string) ([]string, error) {
	ctx := context.Background()
	return p.Pool.BRPop(ctx, timeout, key).Result()
}

// 裁剪列表，保留[start, stop]位置的元素
func (p *RedisProxyer) LTrim(key string, start, stop int64) error {
	ctx := context.Background()
	return p.Pool.LTrim(ctx, key, start, stop).Err()
}
