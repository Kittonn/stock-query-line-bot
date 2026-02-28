package usecases

import (
	"context"
	"log"

	"github.com/Kittonn/stock-query-line-bot/internal/adapters/line_messaging_api/flexmessage/cards"
	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
)

type LineWebhookUsecase struct {
	stockUC          ports.StockUsecase
	lineMessagingAPI ports.LineMessagingAPI
}

func NewLineWebhookUsecase(stockUC ports.StockUsecase, lineMessagingAPI ports.LineMessagingAPI) ports.LineWebhookUsecase {
	return &LineWebhookUsecase{
		stockUC:          stockUC,
		lineMessagingAPI: lineMessagingAPI,
	}
}

func (uc *LineWebhookUsecase) HandleEvent(ctx context.Context, event *domain.LineEvent) {
	switch event.Type {
	case domain.LineEventTypeMessage:
		uc.handleMessage(ctx, event)
	case domain.LineEventTypePostback:
		uc.handlePostback(ctx, event)
	default:
		// ignore other event types for now
	}
}

func (uc *LineWebhookUsecase) handleMessage(ctx context.Context, event *domain.LineEvent) {
	symbol := event.Message.Text
	stock, err := uc.stockUC.GetStockSummary(ctx, symbol)
	if err != nil {
		log.Printf("Error getting stock price for symbol %s: %v", symbol, err)
		return
	}

	log.Printf("stock: %+v", stock)

	card := cards.BuildStockCard(stock)

	err = uc.lineMessagingAPI.Reply(ctx, event.ReplyToken, []domain.Message{
		domain.NewFlexMessage(stock.Name+" Stock Summary", card),
	})

	if err != nil {
		log.Printf("Error replying to LINE message: %v", err)
	}

	log.Printf("Replied with stock summary for %s", symbol)
}

func (uc *LineWebhookUsecase) handlePostback(ctx context.Context, event *domain.LineEvent) {
	// handle postback events if needed
}
