package usecases

import (
	"context"
	"fmt"

	"golang.org/x/sync/singleflight"

	"strings"
	"time"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
	"github.com/Kittonn/stock-query-line-bot/pkg/logger"
)

type StockUsecase struct {
	finnhubAPI ports.FinnHub
	cache      ports.StockCache
	sf         *singleflight.Group
	log        logger.Logger
}

func NewStockUsecase(finnhubAPI ports.FinnHub, cache ports.StockCache, log logger.Logger) ports.StockUsecase {
	return &StockUsecase{
		finnhubAPI: finnhubAPI,
		cache:      cache,
		sf:         &singleflight.Group{},
		log:        log,
	}
}

func (s *StockUsecase) GetStockSummary(ctx context.Context, symbol string) (*domain.StockSummary, error) {
	normalizedSymbol := s.normalizeSymbol(symbol)

	// Use singleflight to prevent multiple concurrent requests for the same symbol
	result, err, shared := s.sf.Do(normalizedSymbol, func() (interface{}, error) {
		stockPrice, err := s.getStockPrice(ctx, normalizedSymbol)
		if err != nil {
			return nil, fmt.Errorf("failed to get stock price: %w", err)
		}

		companyProfile, err := s.getCompanyProfile(ctx, normalizedSymbol)
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

			Name:     companyProfile.Name,
			Exchange: companyProfile.Exchange,
			Ticker:   companyProfile.Ticker,
			Currency: companyProfile.Currency,
		}, nil
	})

	if err != nil {
		s.log.Error("failed to get stock summary symbol: ", normalizedSymbol, " error: ", err)

		return nil, err
	}

	if shared {
		s.log.Info("singleflight shared result for symbol: ", normalizedSymbol)
	}

	return result.(*domain.StockSummary), nil
}

func (s *StockUsecase) getStockPrice(ctx context.Context, symbol string) (*domain.StockPrice, error) {
	// Try to get stock price from cache first
	cachedPrice, err := s.cache.GetStockPrice(ctx, symbol)
	if err == nil && cachedPrice != nil {
		s.log.Info("stock price cache hit symbol: ", symbol)
		return cachedPrice, nil
	}

	// If not in cache, fetch from FinnHub API
	s.log.Info("stock price cache miss symbol: ", symbol)

	quote, err := s.finnhubAPI.GetStockPrice(ctx, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stock price: %w", err)
	}

	// Save the fetched stock price to cache with a short expiration time since stock prices change frequently
	_ = s.cache.SetStockPrice(ctx, symbol, quote, 1*time.Minute)

	return quote, nil
}

func (s *StockUsecase) getCompanyProfile(ctx context.Context, symbol string) (*domain.CompanyProfile, error) {
	// Try to get company profile from cache first
	cachedProfile, err := s.cache.GetCompanyProfile(ctx, symbol)
	if err == nil && cachedProfile != nil {
		s.log.Info("company profile cache hit symbol: ", symbol)
		return cachedProfile, nil
	}

	// If not in cache, fetch from FinnHub API
	s.log.Info("company profile cache miss symbol: ", symbol)
	profile, err := s.finnhubAPI.GetCompanyProfile(ctx, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch company profile: %w", err)
	}

	// Save the fetched company profile to cache with a longer expiration time
	_ = s.cache.SetCompanyProfile(ctx, symbol, profile, 30*24*time.Hour)

	return profile, nil
}

func (s *StockUsecase) normalizeSymbol(symbol string) string {
	return strings.ToUpper(strings.TrimSpace(symbol))
}
