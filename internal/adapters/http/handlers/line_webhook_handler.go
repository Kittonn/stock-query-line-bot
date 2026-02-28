package handlers

import (
	"log"
	"net/http"

	"github.com/Kittonn/stock-query-line-bot/internal/adapters/http/handlers/dto"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
	"github.com/labstack/echo/v5"
)

type LineWebhookHandler struct {
	lineWebhookWorker ports.LineWebhookWorker
}

func NewLineWebhookHandler(lineWebhookWorker ports.LineWebhookWorker) *LineWebhookHandler {
	return &LineWebhookHandler{
		lineWebhookWorker: lineWebhookWorker,
	}
}

func (h *LineWebhookHandler) Handle(c *echo.Context) error {
	webhookEvent := new(dto.WebhookRequest)
	if err := c.Bind(webhookEvent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	go h.processEvents(webhookEvent.Events)

	return c.NoContent(http.StatusNoContent)
}

func (h *LineWebhookHandler) processEvents(events []dto.Event) {
	for _, event := range events {
		domainEvent, err := event.ToDomain()
		if err != nil {
			continue
		}

		log.Printf("Received event: %s %s\n", domainEvent.Type, domainEvent.Message)
		h.lineWebhookWorker.Enqueue(domainEvent)
	}
}
