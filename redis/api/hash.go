package api

import (
	"context"
)

/*
*
  - @Description: 同时设置一个或多个 key-value 对
  - @receiver p
  - @param key
  - @param kvs: HSet accepts values in following formats:
  - HSet("myhash", "key1", "value1", "key2", "value2")
  - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
  - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
  - @return error
*/
func (p *RedisProxyer) HSet(key string, kvs ...interface{}) error {
	ctx := context.Background()
	return p.Pool.HSet(ctx, key, kvs...).Err()
}

// 查看哈希表 key 中，指定的字段是否存在
func (p *RedisProxyer) HExists(key, field string) (bool, error) {
	ctx := context.Background()
	return p.Pool.HExists(ctx, key, field).Result()
}

// 删除一个或多个哈希表字段
func (p *RedisProxyer) HDel(key string, fields ...string) error {
	ctx := context.Background()
	return p.Pool.HDel(ctx, key, fields...).Err()
}

// 获取在哈希表中指定 key 的所有字段和值
func (p *RedisProxyer) HGetAll(key string) (map[string]string, error) {
	ctx := context.Background()
	return p.Pool.HGetAll(ctx, key).Result()
}

// 获取存储在哈希表中指定字段的值
func (p *RedisProxyer) HGet(key, field string) (string, error) {
	ctx := context.Background()
	return p.Pool.HGet(ctx, key, field).Result()
}

func (p *RedisProxyer) HGetBytes(key, field string) ([]byte, error) {
	ctx := context.Background()
	return p.Pool.HGet(ctx, key, field).Bytes()
}

func (p *RedisProxyer) HGetBool(key, field string) (bool, error) {
	ctx := context.Background()
	return p.Pool.HGet(ctx, key, field).Bool()
}

func (p *RedisProxyer) HGetFloat64(key, field string) (float64, error) {
	ctx := context.Background()
	return p.Pool.HGet(ctx, key, field).Float64()
}

func (p *RedisProxyer) HGetFloat32(key, field string) (float32, error) {
	ctx := context.Background()
	return p.Pool.HGet(ctx, key, field).Float32()
}

func (p *RedisProxyer) HGetInt(key, field string) (int, error) {
	ctx := context.Background()
	return p.Pool.HGet(ctx, key, field).Int()
}

func (p *RedisProxyer) HGetInt64(key, field string) (int64, error) {
	ctx := context.Background()
	return p.Pool.HGet(ctx, key, field).Int64()
}

func (p *RedisProxyer) HGetUint64(key, field string) (uint64, error) {
	ctx := context.Background()
	return p.Pool.HGet(ctx, key, field).Uint64()
}

// 为哈希表 key 中的指定字段的整数值加上增量 increment
func (p *RedisProxyer) HIncrBy(key, field string, incr int64) (int64, error) {
	ctx := context.Background()
	return p.Pool.HIncrBy(ctx, key, field, incr).Result()
}
