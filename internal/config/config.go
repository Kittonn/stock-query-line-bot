package config

import (
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

type Config struct {
	Port int `mapstructure:"PORT"`

	// Circuit Breaker Settings
	CircuitBreakerOpenStateTimeout                 time.Duration `mapstructure:"CIRCUIT_BREAKER_OPEN_STATE_TIMEOUT"`
	CircuitBreakerHalfOpenStateMaximumRequestCount uint32        `mapstructure:"CIRCUIT_BREAKER_HALF_OPEN_STATE_MAXIMUM_REQUEST_COUNT"`
	CircuitBreakerConsecutiveFailureThreshold      uint32        `mapstructure:"CIRCUIT_BREAKER_CONSECUTIVE_FAILURE_THRESHOLD"`
	CircuitBreakerBucketPeriod                     time.Duration `mapstructure:"CIRCUIT_BREAKER_BUCKET_PERIOD"`

	// Finnhub API Settings
	FinnhubAPIKey string `mapstructure:"FINNHUB_API_KEY"`
	FinnhubAPIURL string `mapstructure:"FINNHUB_API_URL"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	_ = v.ReadInConfig()

	setDefaultConfig(v)

	var config Config
	if err := v.Unmarshal(&config, func(dc *mapstructure.DecoderConfig) {
		dc.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
		)
	}); err != nil {
		return nil, err
	}

	return &config, nil
}

func setDefaultConfig(v *viper.Viper) {
	v.SetDefault("PORT", 8080)

	// Circuit Breaker default settings
	v.SetDefault("CIRCUIT_BREAKER_OPEN_STATE_TIMEOUT", "60s")
	v.SetDefault("CIRCUIT_BREAKER_HALF_OPEN_STATE_MAXIMUM_REQUEST_COUNT", 3)
	v.SetDefault("CIRCUIT_BREAKER_CONSECUTIVE_FAILURE_THRESHOLD", 5)
	v.SetDefault("CIRCUIT_BREAKER_BUCKET_PERIOD", "10s")
}
