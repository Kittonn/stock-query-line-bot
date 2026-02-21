package finnhubapi

import (
	"context"

	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"github.com/go-resty/resty/v2"
	"github.com/sony/gobreaker/v2"
)

type Quote struct {
	CurrentPrice        float64 `json:"c"`  // Current price
	PriceChange         float64 `json:"d"`  // Change
	PercentChange       float64 `json:"dp"` // Percent change
	HighPriceOfDay      float64 `json:"h"`  // High price of the day
	LowPriceOfDay       float64 `json:"l"`  // Low price of the day
	OpenPriceOfDay      float64 `json:"o"`  // Open price of the day
	PreviousClosePrice  float64 `json:"pc"` // Previous close price
	TimestampUnixSecond int64   `json:"t"`  // Timestamp
}

type FinnhubAPI struct {
	client  *resty.Client
	breaker *gobreaker.CircuitBreaker[*Quote]
	cfg     *config.Config
}

func NewFinnhubAPI(cfg *config.Config, client *resty.Client) *FinnhubAPI {

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

func (f *FinnhubAPI) GetQuote(ctx context.Context, symbol string) (*Quote, error) {
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
			return nil, err
		}

		if resp.IsError() {
			return nil, err
		}

		return resp.Result().(*Quote), nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
