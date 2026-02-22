package ports

import (
	"context"
	"time"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
)

type StockCache interface {
	GetStockPrice(ctx context.Context, symbol string) (*domain.StockPrice, error)
	SetStockPrice(ctx context.Context, symbol string, price *domain.StockPrice, ttl time.Duration) error

	GetCompanyProfile(ctx context.Context, symbol string) (*domain.CompanyProfile, error)
	SetCompanyProfile(ctx context.Context, symbol string, profile *domain.CompanyProfile, ttl time.Duration) error
}
