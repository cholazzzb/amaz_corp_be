package database

import "github.com/redis/go-redis/v9"

type RedisRepository struct {
	Rds *redis.Client
}

func NewRedisRepository(rds *redis.Client) *RedisRepository {
	return &RedisRepository{
		Rds: rds,
	}
}
