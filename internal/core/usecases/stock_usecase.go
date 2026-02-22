package usecases

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
)

type StockUsecase struct {
	finnhubAPI ports.FinnHub
	cache      ports.StockCache
}

func NewStockUsecase(finnhubAPI ports.FinnHub, cache ports.StockCache) ports.StockUsecase {
	return &StockUsecase{
		finnhubAPI: finnhubAPI,
		cache:      cache,
	}
}

func (s *StockUsecase) GetStockPrice(ctx context.Context, symbol string) (*domain.StockPrice, error) {
	normalizedSymbol := s.normalizeSymbol(symbol)
	cachedPrice, err := s.cache.GetStockPrice(ctx, normalizedSymbol)
	if err == nil && cachedPrice != nil {
		return cachedPrice, nil
	}

	quote, err := s.finnhubAPI.GetStockPrice(ctx, normalizedSymbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stock price: %w", err)
	}

	// Stock price is changed often, so we only cache it for a short time
	_ = s.cache.SetStockPrice(ctx, normalizedSymbol, quote, 1*time.Minute)

	return quote, nil
}

func (s *StockUsecase) GetCompanyProfile(ctx context.Context, symbol string) (*domain.CompanyProfile, error) {
	normalizedSymbol := s.normalizeSymbol(symbol)
	cachedProfile, err := s.cache.GetCompanyProfile(ctx, normalizedSymbol)
	if err == nil && cachedProfile != nil {
		return cachedProfile, nil
	}

	profile, err := s.finnhubAPI.GetCompanyProfile(ctx, normalizedSymbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch company profile: %w", err)
	}

	// Company profile is not changed often, so we can cache it for a long time
	_ = s.cache.SetCompanyProfile(ctx, normalizedSymbol, profile, 30*24*time.Hour)

	return profile, nil
}

func (s *StockUsecase) normalizeSymbol(symbol string) string {
	return strings.ToUpper(strings.TrimSpace(symbol))
}
