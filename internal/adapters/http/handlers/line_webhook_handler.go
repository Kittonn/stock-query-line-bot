package handlers

import (
	"net/http"

	"github.com/Kittonn/stock-query-line-bot/internal/adapters/http/handlers/dto"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
	"github.com/Kittonn/stock-query-line-bot/pkg/logger"
	"github.com/labstack/echo/v5"
)

type LineWebhookHandler struct {
	lineWebhookWorker ports.LineWebhookWorker
	log               logger.Logger
}

func NewLineWebhookHandler(lineWebhookWorker ports.LineWebhookWorker, log logger.Logger) *LineWebhookHandler {
	return &LineWebhookHandler{
		lineWebhookWorker: lineWebhookWorker,
		log:               log,
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
			h.log.Warn("failed to parse event: ", err)
			continue
		}

		h.log.Info("received event: ", domainEvent.Type)
		h.lineWebhookWorker.Enqueue(domainEvent)
	}
}
