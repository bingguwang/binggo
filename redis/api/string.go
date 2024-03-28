package api

import (
	"context"
	"time"
)

// 设置指定 key 的值
func (p *RedisProxyer) Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	return p.Pool.Set(ctx, key, value, expiration).Err()
}

/*
*
  - @Description: 同时设置一个或多个 key-value 对
  - @receiver p
  - @param key
  - @param kvs: MSet is like Set but accepts multiple values:
  - MSet("key1", "value1", "key2", "value2")
  - MSet([]string{"key1", "value1", "key2", "value2"})
  - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
  - @return error
*/
func (p *RedisProxyer) MSet(kvs ...interface{}) error {
	ctx := context.Background()
	return p.Pool.MSet(ctx, kvs...).Err()
}

/**
 * @Description: 检查key是否存在
 * @receiver p
 * @param key: 要检查的key
 * @return int64: 若 key 存在返回 1 ，否则返回 0
 * @return error
 */
func (p *RedisProxyer) Exists(key string) (int64, error) {
	ctx := context.Background()
	return p.Pool.Exists(ctx, key).Result()
}

// 删除键值
func (p *RedisProxyer) Del(key ...string) error {
	ctx := context.Background()
	return p.Pool.Del(ctx, key...).Err()
}

// 获取指定键值
func (p *RedisProxyer) Get(key string) (string, error) {
	ctx := context.Background()
	return p.Pool.Get(ctx, key).Result()
}

func (p *RedisProxyer) MGet(keys ...string) ([]interface{}, error) {
	ctx := context.Background()
	return p.Pool.MGet(ctx, keys...).Result()
}

func (p *RedisProxyer) GetBytes(key string) ([]byte, error) {
	ctx := context.Background()
	return p.Pool.Get(ctx, key).Bytes()
}

func (p *RedisProxyer) GetBool(key string) (bool, error) {
	ctx := context.Background()
	return p.Pool.Get(ctx, key).Bool()
}

func (p *RedisProxyer) GetFloat64(key string) (float64, error) {
	ctx := context.Background()
	return p.Pool.Get(ctx, key).Float64()
}

func (p *RedisProxyer) GetFloat32(key string) (float32, error) {
	ctx := context.Background()
	return p.Pool.Get(ctx, key).Float32()
}

func (p *RedisProxyer) GetInt(key string) (int, error) {
	ctx := context.Background()
	return p.Pool.Get(ctx, key).Int()
}

func (p *RedisProxyer) GetInt64(key string) (int64, error) {
	ctx := context.Background()
	return p.Pool.Get(ctx, key).Int64()
}

func (p *RedisProxyer) GetUint64(key string) (uint64, error) {
	ctx := context.Background()
	return p.Pool.Get(ctx, key).Uint64()
}

// 将 key 中储存的数字值增一
func (p *RedisProxyer) Incr(key string) (int64, error) {
	ctx := context.Background()
	return p.Pool.Incr(ctx, key).Result()
}

// 将 key 所储存的值加上给定的增量值
func (p *RedisProxyer) IncrBy(key string, value int64) (int64, error) {
	ctx := context.Background()
	return p.Pool.IncrBy(ctx, key, value).Result()
}

// 将 key 中储存的数字值减一
func (p *RedisProxyer) Decr(key string) (int64, error) {
	ctx := context.Background()
	return p.Pool.Decr(ctx, key).Result()
}

// 将 key 所储存的值减去给定的减量值
func (p *RedisProxyer) DecrBy(key string, value int64) (int64, error) {
	ctx := context.Background()
	return p.Pool.DecrBy(ctx, key, value).Result()
}
