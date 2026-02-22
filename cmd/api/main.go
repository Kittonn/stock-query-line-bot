package main

import (
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/cache"
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/finnhub_api"
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/http"
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/http_client"
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/line_api"
	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"github.com/Kittonn/stock-query-line-bot/internal/core/usecases"
	"github.com/Kittonn/stock-query-line-bot/internal/infrastructure"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		http_client.Module,
		finnhub_api.Module,
		line_api.Module,
		http.Module,
		infrastructure.Module,
		cache.Module,
		usecases.Module,
	).Run()
}
