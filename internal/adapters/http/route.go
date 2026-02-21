package http

import (
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/http/middlewares"
	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(cfg *config.Config, e *echo.Echo, h *Handlers) {
	e.GET("/health", h.Health.HealthCheck)

	v1 := e.Group("/api/v1")
	v1.POST("/line-webhook", h.LineWebhook.Handle, middlewares.VerifyLineSignature(cfg.LineChannelSecret))
}
