package main

import (
	finnhubapi "github.com/Kittonn/stock-query-line-bot/internal/adapters/finnhub_api"
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/http"
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/http_client"
	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		http_client.Module,
		finnhubapi.Module,
		http.Module,
	).Run()
}
