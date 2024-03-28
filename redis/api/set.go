package api

import (
	"context"
)

// 向集合添加一个或多个成员
func (p *RedisProxyer) SAdd(key string, members ...interface{}) error {
	ctx := context.Background()
	return p.Pool.SAdd(ctx, key, members...).Err()
}

// 返回集合中的所有成员
func (p *RedisProxyer) SMembers(key string) ([]string, error) {
	ctx := context.Background()
	return p.Pool.SMembers(ctx, key).Result()
}

// 判断 member 元素是否是集合 key 的成员
func (p *RedisProxyer) SIsMember(key string, member interface{}) (bool, error) {
	ctx := context.Background()
	return p.Pool.SIsMember(ctx, key, member).Result()
}

// 移除集合中一个或多个成员
func (p *RedisProxyer) SRem(key string, members ...interface{}) error {
	ctx := context.Background()
	return p.Pool.SRem(ctx, key, members...).Err()
}

// 返回给定所有集合的交集
func (p *RedisProxyer) SInter(keys ...string) ([]string, error) {
	ctx := context.Background()
	return p.Pool.SInter(ctx, keys...).Result()
}

// 返回第一个集合与其他集合之间的差异
func (p *RedisProxyer) SDiff(keys ...string) ([]string, error) {
	ctx := context.Background()
	return p.Pool.SDiff(ctx, keys...).Result()
}

// 返回所有给定集合的并集
func (p *RedisProxyer) SUnion(keys ...string) ([]string, error) {
	ctx := context.Background()
	return p.Pool.SUnion(ctx, keys...).Result()
}
