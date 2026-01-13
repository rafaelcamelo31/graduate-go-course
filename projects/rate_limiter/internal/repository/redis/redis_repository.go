package repository

type RedisRepository struct{}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{}
}
