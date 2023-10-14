package cache

import (
	redis "github.com/redis/go-redis/v9"
)

var RedisDB *redis.Client

func InitRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "tiktok",
		DB:       4,
	})
	if RedisDB == nil {
		panic("[redis] init error")
	}
}
