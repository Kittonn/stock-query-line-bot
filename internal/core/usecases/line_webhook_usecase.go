package usecases

import (
	"context"

	"github.com/Kittonn/stock-query-line-bot/internal/adapters/line_messaging_api/flexmessage/cards"
	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
	"github.com/Kittonn/stock-query-line-bot/pkg/logger"
)

type LineWebhookUsecase struct {
	stockUC          ports.StockUsecase
	lineMessagingAPI ports.LineMessagingAPI
	log              logger.Logger
}

func NewLineWebhookUsecase(stockUC ports.StockUsecase, lineMessagingAPI ports.LineMessagingAPI, log logger.Logger) ports.LineWebhookUsecase {
	return &LineWebhookUsecase{
		stockUC:          stockUC,
		lineMessagingAPI: lineMessagingAPI,
		log:              log,
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
		uc.log.Error("failed to get stock summary symbol: ", symbol, " error: ", err)
		return
	}

	card := cards.BuildStockCard(stock)

	err = uc.lineMessagingAPI.Reply(ctx, event.ReplyToken, []domain.Message{
		domain.NewFlexMessage(stock.Name+" Stock Summary", card),
	})

	if err != nil {
		uc.log.Error("failed to reply to LINE message symbol: ", symbol, " error: ", err)
		return
	}

	uc.log.Info("replied with stock summary symbol: ", symbol)

}

func (uc *LineWebhookUsecase) handlePostback(ctx context.Context, event *domain.LineEvent) {
	// handle postback events if needed
}
