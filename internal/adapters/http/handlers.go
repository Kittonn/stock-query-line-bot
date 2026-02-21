package http

import "github.com/Kittonn/stock-query-line-bot/internal/adapters/http/handlers"

type Handlers struct {
	Health      *handlers.HealthHandler
	LineWebhook *handlers.LineWebhookHandler
}

func NewHandlers(
	health *handlers.HealthHandler,
	lineWebhook *handlers.LineWebhookHandler,
) *Handlers {
	return &Handlers{
		Health:      health,
		LineWebhook: lineWebhook,
	}
}
