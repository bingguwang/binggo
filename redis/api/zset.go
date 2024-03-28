package api

import (
	"context"

	"github.com/go-redis/redis/v8"
)

/**
 * @Description:向有序集合添加一个或多个成员，或者更新已存在成员的分数
 * @receiver p
 * @param key: 有序集合键
 * @param members: 成员和分数的键值对
 * @return error
 */
func (p *RedisProxyer) ZAdd(key string, members map[interface{}]float64) error {
	ctx := context.Background()
	zsetMembers := make([]*redis.Z, 0)
	for member, score := range members {
		z := &redis.Z{
			Score:  score,
			Member: member,
		}
		zsetMembers = append(zsetMembers, z)
	}
	return p.Pool.ZAdd(ctx, key, zsetMembers...).Err()
}

// 移除有序集合中的一个或多个成员
func (p *RedisProxyer) ZRem(key string, members ...interface{}) error {
	ctx := context.Background()
	return p.Pool.ZRem(ctx, key, members...).Err()
}

/**
 * @Description:通过分数返回有序集合指定区间内的成员
 * @receiver p
 * @param key
 * @param min:
 * @param max:
 * @return []string
 * @return error
 */
func (p *RedisProxyer) ZRangeByScore(key, min, max string) ([]string, error) {
	ctx := context.Background()
	return p.Pool.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min: min,
		Max: max,
	}).Result()
}

/**
 * @Description:通过分数返回有序集合指定区间内的成员,带有Score
 * @receiver p
 * @param key
 * @param min:
 * @param max:
 * @return []float64 score列表
 * @return []interface{} member列表
 * @return error
 */
func (p *RedisProxyer) ZRangeByScoreWithScores(key, min, max string) ([]float64, []interface{}, error) {
	ctx := context.Background()
	set, err := p.Pool.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min: min,
		Max: max,
	}).Result()
	if err != nil {
		return nil, nil, err
	}
	scores := make([]float64, len(set))
	members := make([]interface{}, len(set))
	for i := 0; i < len(set); i++ {
		scores[i] = set[i].Score
		members[i] = set[i].Member
	}
	return scores, members, nil
}

/**
 * @Description: 移除有序集合中给定的分数区间的所有成员
 * @receiver p
 * @param key
 * @param min:
 * @param max:
 * @return []string
 * @return error
 */
func (p *RedisProxyer) ZRemRangeByScore(key, min, max string) error {
	ctx := context.Background()
	return p.Pool.ZRemRangeByScore(ctx, key, min, max).Err()
}
