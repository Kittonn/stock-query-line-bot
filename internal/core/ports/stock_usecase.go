package ports

import (
	"context"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
)

type StockUsecase interface {
	GetStockSummary(ctx context.Context, symbol string) (*domain.StockSummary, error)
}
