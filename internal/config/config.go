package config

import (
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

type Config struct {
	// Server Settings
	Port       int    `mapstructure:"PORT"`
	AppEnv     string `mapstructure:"APP_ENV"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	LogEncoder string `mapstructure:"LOG_ENCODER"`

	// Circuit Breaker Settings
	CircuitBreakerOpenStateTimeout                 time.Duration `mapstructure:"CIRCUIT_BREAKER_OPEN_STATE_TIMEOUT"`
	CircuitBreakerHalfOpenStateMaximumRequestCount uint32        `mapstructure:"CIRCUIT_BREAKER_HALF_OPEN_STATE_MAXIMUM_REQUEST_COUNT"`
	CircuitBreakerConsecutiveFailureThreshold      uint32        `mapstructure:"CIRCUIT_BREAKER_CONSECUTIVE_FAILURE_THRESHOLD"`
	CircuitBreakerBucketPeriod                     time.Duration `mapstructure:"CIRCUIT_BREAKER_BUCKET_PERIOD"`

	// Finnhub API Settings
	FinnhubAPIKey string `mapstructure:"FINNHUB_API_KEY"`
	FinnhubAPIURL string `mapstructure:"FINNHUB_API_BASE_URL"`

	// Line Messaging API Settings
	LineChannelSecret      string `mapstructure:"LINE_CHANNEL_SECRET"`
	LineChannelAccessToken string `mapstructure:"LINE_CHANNEL_ACCESS_TOKEN"`
	LineMessagingAPIURL    string `mapstructure:"LINE_MESSAGING_API_BASE_URL"`

	// Redis Settings
	RedisAddr            string        `mapstructure:"REDIS_ADDR"`
	RedisPassword        string        `mapstructure:"REDIS_PASSWORD"`
	RedisDB              int           `mapstructure:"REDIS_DB"`
	RedisPoolSize        int           `mapstructure:"REDIS_POOL_SIZE"`
	RedisPoolTimeout     time.Duration `mapstructure:"REDIS_POOL_TIMEOUT"`
	RedisMinIdleConns    int           `mapstructure:"REDIS_MIN_IDLE_CONNS"`
	RedisConnMaxIdleTime time.Duration `mapstructure:"REDIS_CONN_MAX_IDLE_TIME"`
	RedisConnMaxLifetime time.Duration `mapstructure:"REDIS_CONN_MAX_LIFETIME"`
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
	// Server default settings
	v.SetDefault("PORT", 8080)
	v.SetDefault("APP_ENV", "development")
	v.SetDefault("LOG_LEVEL", "info")
	v.SetDefault("LOG_ENCODER", "console")

	// Circuit Breaker default settings
	v.SetDefault("CIRCUIT_BREAKER_OPEN_STATE_TIMEOUT", "60s")
	v.SetDefault("CIRCUIT_BREAKER_HALF_OPEN_STATE_MAXIMUM_REQUEST_COUNT", 3)
	v.SetDefault("CIRCUIT_BREAKER_CONSECUTIVE_FAILURE_THRESHOLD", 5)
	v.SetDefault("CIRCUIT_BREAKER_BUCKET_PERIOD", "10s")

	// Finnhub API default settings
	v.SetDefault("FINNHUB_API_BASE_URL", "https://finnhub.io/api/v1")

	// Line Messaging API default settings
	v.SetDefault("LINE_MESSAGING_API_BASE_URL", "https://api.line.me/v2/bot/message")

	// Redis default settings
	v.SetDefault("REDIS_ADDR", "localhost:6379")
	v.SetDefault("REDIS_PASSWORD", "")
	v.SetDefault("REDIS_DB", 0)
	v.SetDefault("REDIS_POOL_SIZE", 30)
	v.SetDefault("REDIS_POOL_TIMEOUT", "5m")
	v.SetDefault("REDIS_MIN_IDLE_CONNS", 5)
	v.SetDefault("REDIS_CONN_MAX_IDLE_TIME", "1m")
	v.SetDefault("REDIS_CONN_MAX_LIFETIME", "5m")
}
