package redis_storage

import "github.com/go-redis/redis"

type Config struct {
	Addr string
}

func NewRedisClient(cfg Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: "",
	})
	return rdb, nil
}
