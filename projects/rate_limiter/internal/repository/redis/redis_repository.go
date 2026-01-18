package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: client,
	}
}

func (r *RedisRepository) IsBlocked(ctx context.Context, key string) (bool, error) {
	exists, err := r.client.Exists(ctx, blockKey(key)).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (r *RedisRepository) GetCount(ctx context.Context, key string) (int64, error) {
	return r.client.Get(ctx, countKey(key)).Int64()
}

func (r *RedisRepository) Increment(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, countKey(key)).Result()
}

func (r *RedisRepository) SetCount(ctx context.Context, key string, window time.Duration) error {
	return r.client.Expire(ctx, countKey(key), window).Err()
}

func (r *RedisRepository) Block(ctx context.Context, key string, blockDuration time.Duration) error {
	return r.client.Set(ctx, blockKey(key), "1", blockDuration).Err()
}

func countKey(key string) string {
	return "rl:" + key + ":count"
}

func blockKey(key string) string {
	return "rl:" + key + ":block"
}
