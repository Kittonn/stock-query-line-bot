package usecases

import (
	"context"
	"fmt"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
)

type StockUsecase struct {
	finnhubAPI ports.FinnHub
}

func NewStockUsecase(finnhubAPI ports.FinnHub) ports.StockUsecase {
	return &StockUsecase{
		finnhubAPI: finnhubAPI,
	}
}

func (s *StockUsecase) GetStockPrice(ctx context.Context, symbol string) (*domain.StockPrice, error) {
	// TODO: Add caching logic here to reduce the number of API calls to FinnHub

	quote, err := s.finnhubAPI.GetStockPrice(ctx, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stock price: %w", err)
	}

	return quote, nil
}
