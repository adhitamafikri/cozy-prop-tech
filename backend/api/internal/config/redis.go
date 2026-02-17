package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})

	return rdb
}

type RedisConfig struct {
	Host   string
	Port   int
	Prefix string
}

func (c *RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
