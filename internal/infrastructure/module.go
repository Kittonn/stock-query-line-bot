package infrastructure

import (
	"github.com/Kittonn/stock-query-line-bot/internal/infrastructure/redis"
	"go.uber.org/fx"
)

var Module = fx.Module("infrastructure", fx.Provide(redis.NewRedisClient))
