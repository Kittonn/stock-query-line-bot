package main

import (
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/http"
	"github.com/Kittonn/stock-query-line-bot/internal/app"
	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		app.Module,
		http.Module,
	).Run()
}
