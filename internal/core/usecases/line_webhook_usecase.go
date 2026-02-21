package usecases

import (
	"context"
	"log"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
)

type LineWebhookUsecase struct {
	stockUC ports.StockUsecase
}

func NewLineWebhookUsecase(stockUC ports.StockUsecase) ports.LineWebhookUsecase {
	return &LineWebhookUsecase{
		stockUC: stockUC,
	}
}

func (uc *LineWebhookUsecase) HandleEvent(ctx context.Context, event domain.LineEvent) {
	switch event.Type {
	case domain.LineEventTypeMessage:
		log.Printf("Handling message event: %s", event.Message)
		symbol := event.Message
		stockPrice, err := uc.stockUC.GetStockPrice(ctx, symbol)
		if err != nil {
			log.Printf("Error getting stock price for symbol %s: %v", symbol, err)
			return
		}

		// TODO: Send the stock price back to the user via LINE Messaging API

		log.Printf("Stock price for %s: Current: %.2f, High: %.2f, Low: %.2f\n", symbol, stockPrice.CurrentPrice, stockPrice.HighPriceOfDay, stockPrice.LowPriceOfDay)
	default:
		// ignore other event types for now
	}
}
