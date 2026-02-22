package finnhub_api

import (
	"context"
	"log"

	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
	"github.com/go-resty/resty/v2"
	"github.com/sony/gobreaker/v2"
)

type FinnhubAPI struct {
	client  *resty.Client
	breaker *gobreaker.CircuitBreaker[*Quote]
	cfg     *config.Config
}

func NewFinnhubAPI(cfg *config.Config, client *resty.Client) ports.FinnHub {

	return &FinnhubAPI{
		cfg:    cfg,
		client: client,
		breaker: gobreaker.NewCircuitBreaker[*Quote](gobreaker.Settings{
			Name:         "finnhub",
			Timeout:      cfg.CircuitBreakerOpenStateTimeout,
			BucketPeriod: cfg.CircuitBreakerBucketPeriod,
			MaxRequests:  2,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				return counts.ConsecutiveFailures > cfg.CircuitBreakerConsecutiveFailureThreshold
			},
		}),
	}
}

func (f *FinnhubAPI) GetStockPrice(ctx context.Context, symbol string) (*domain.StockPrice, error) {
	log.Printf("Attempting to get stock price for symbol: %s", symbol)

	result, err := f.breaker.Execute(func() (*Quote, error) {
		resp, err := f.client.R().
			SetContext(ctx).
			SetQueryParams(map[string]string{
				"symbol": symbol,
			}).
			SetHeader("X-Finnhub-Token", f.cfg.FinnhubAPIKey).
			SetResult(&Quote{}).
			Get(f.cfg.FinnhubAPIURL + "/quote")

		if err != nil {
			log.Printf("FinnHub HTTP request error: %v", err)
			return nil, err
		}

		if resp.IsError() {
			log.Printf("FinnHub HTTP response error. Status: %d, Body: %s", resp.StatusCode(), resp.String())
			return nil, err
		}

		return resp.Result().(*Quote), nil
	})

	if err != nil {
		return nil, err
	}

	return &domain.StockPrice{
		CurrentPrice:       result.CurrentPrice,
		HighPriceOfDay:     result.HighPriceOfDay,
		LowPriceOfDay:      result.LowPriceOfDay,
		OpenPriceOfDay:     result.OpenPriceOfDay,
		PreviousClosePrice: result.PreviousClosePrice,
		PriceChange:        result.PriceChange,
		PercentChange:      result.PercentChange,
	}, nil
}

func (f *FinnhubAPI) GetCompanyProfile(ctx context.Context, symbol string) (*domain.CompanyProfile, error) {
	resp, err := f.client.R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"symbol": symbol,
		}).
		SetHeader("X-Finnhub-Token", f.cfg.FinnhubAPIKey).
		SetResult(&CompanyProfile{}).
		Get(f.cfg.FinnhubAPIURL + "/stock/profile2")

	if err != nil {
		log.Printf("FinnHub HTTP request error: %v", err)
		return nil, err
	}

	if resp.IsError() {
		log.Printf("FinnHub HTTP response error. Status: %d, Body: %s", resp.StatusCode(), resp.String())
		return nil, err
	}

	result := resp.Result().(*CompanyProfile)
	return &domain.CompanyProfile{
		Country:              result.Country,
		Currency:             result.Currency,
		Exchange:             result.Exchange,
		IPO:                  result.IPO,
		MarketCapitalization: result.MarketCapitalization,
		Name:                 result.Name,
		Phone:                result.Phone,
		ShareOutstanding:     result.ShareOutstanding,
		Logo:                 result.Logo,
		FinnhubIndustry:      result.FinnhubIndustry,
	}, nil
}
