package redis

import (
	"context"
	"log"

	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:            cfg.RedisAddr,
		Password:        cfg.RedisPassword,
		DB:              cfg.RedisDB,
		PoolSize:        cfg.RedisPoolSize,
		PoolTimeout:     cfg.RedisPoolTimeout,
		MinIdleConns:    cfg.RedisMinIdleConns,
		ConnMaxIdleTime: cfg.RedisConnMaxIdleTime,
		ConnMaxLifetime: cfg.RedisConnMaxLifetime,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	return client
}
