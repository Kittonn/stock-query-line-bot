package ports

import (
	"context"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
)

type FinnHub interface {
	GetStockPrice(ctx context.Context, symbol string) (*domain.StockPrice, error)
}
