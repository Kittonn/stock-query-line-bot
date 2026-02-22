package usecases

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"

	"strings"
	"time"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
)

type StockUsecase struct {
	finnhubAPI ports.FinnHub
	cache      ports.StockCache
	sf         *singleflight.Group
}

func NewStockUsecase(finnhubAPI ports.FinnHub, cache ports.StockCache) ports.StockUsecase {
	return &StockUsecase{
		finnhubAPI: finnhubAPI,
		cache:      cache,
		sf:         &singleflight.Group{},
	}
}

func (s *StockUsecase) GetStockSummary(ctx context.Context, symbol string) (*domain.StockSummary, error) {
	normalizedSymbol := s.normalizeSymbol(symbol)

	// Use singleflight to prevent multiple concurrent requests for the same symbol
	result, err, _ := s.sf.Do(normalizedSymbol, func() (interface{}, error) {
		stockPrice, err := s.GetStockPrice(ctx, normalizedSymbol)
		if err != nil {
			return nil, fmt.Errorf("failed to get stock price: %w", err)
		}

		companyProfile, err := s.GetCompanyProfile(ctx, normalizedSymbol)
		if err != nil {
			return nil, fmt.Errorf("failed to get company profile: %w", err)
		}

		return &domain.StockSummary{
			CurrentPrice:       stockPrice.CurrentPrice,
			PriceChange:        stockPrice.PriceChange,
			PercentChange:      stockPrice.PercentChange,
			HighPriceOfDay:     stockPrice.HighPriceOfDay,
			LowPriceOfDay:      stockPrice.LowPriceOfDay,
			OpenPriceOfDay:     stockPrice.OpenPriceOfDay,
			PreviousClosePrice: stockPrice.PreviousClosePrice,
			Name:               companyProfile.Name,
			Exchange:           companyProfile.Exchange,
			Ticker:             companyProfile.Ticker,
			Currency:           companyProfile.Currency,
		}, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*domain.StockSummary), nil
}

func (s *StockUsecase) GetStockPrice(ctx context.Context, symbol string) (*domain.StockPrice, error) {
	normalizedSymbol := s.normalizeSymbol(symbol)

	// Try to get stock price from cache first
	cachedPrice, err := s.cache.GetStockPrice(ctx, normalizedSymbol)
	if err == nil && cachedPrice != nil {
		return cachedPrice, nil
	}

	// If not in cache, fetch from FinnHub API
	quote, err := s.finnhubAPI.GetStockPrice(ctx, normalizedSymbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stock price: %w", err)
	}

	// Save the fetched stock price to cache with a short expiration time since stock prices change frequently
	_ = s.cache.SetStockPrice(ctx, normalizedSymbol, quote, 1*time.Minute)

	return quote, nil
}

func (s *StockUsecase) GetCompanyProfile(ctx context.Context, symbol string) (*domain.CompanyProfile, error) {
	normalizedSymbol := s.normalizeSymbol(symbol)

	// Try to get company profile from cache first
	cachedProfile, err := s.cache.GetCompanyProfile(ctx, normalizedSymbol)
	if err == nil && cachedProfile != nil {
		return cachedProfile, nil
	}

	// If not in cache, fetch from FinnHub API
	profile, err := s.finnhubAPI.GetCompanyProfile(ctx, normalizedSymbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch company profile: %w", err)
	}

	// Save the fetched company profile to cache with a longer expiration time
	_ = s.cache.SetCompanyProfile(ctx, normalizedSymbol, profile, 30*24*time.Hour)

	return profile, nil
}

func (s *StockUsecase) normalizeSymbol(symbol string) string {
	return strings.ToUpper(strings.TrimSpace(symbol))
}
