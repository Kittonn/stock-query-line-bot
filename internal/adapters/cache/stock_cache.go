package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
	"github.com/redis/go-redis/v9"
)

const (
	stockPriceKey   = "stock:price:%s"
	stockProfileKey = "stock:profile:%s"
)

type redisCache struct {
	client *redis.Client
}

func NewStockCache(client *redis.Client) ports.StockCache {
	return &redisCache{
		client: client,
	}
}

func (r *redisCache) GetStockPrice(ctx context.Context, symbol string) (*domain.StockPrice, error) {
	key := fmt.Sprintf(stockPriceKey, symbol)
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var price domain.StockPrice
	if err := json.Unmarshal([]byte(val), &price); err != nil {
		return nil, err
	}

	return &price, nil
}

func (r *redisCache) SetStockPrice(ctx context.Context, symbol string, price *domain.StockPrice, ttl time.Duration) error {
	key := fmt.Sprintf(stockPriceKey, symbol)

	data, err := json.Marshal(price)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *redisCache) GetCompanyProfile(ctx context.Context, symbol string) (*domain.CompanyProfile, error) {
	key := fmt.Sprintf(stockProfileKey, symbol)
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var profile domain.CompanyProfile
	if err := json.Unmarshal([]byte(val), &profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *redisCache) SetCompanyProfile(ctx context.Context, symbol string, profile *domain.CompanyProfile, ttl time.Duration) error {
	key := fmt.Sprintf(stockProfileKey, symbol)
	data, err := json.Marshal(profile)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, ttl).Err()
}
