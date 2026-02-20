package http

import (
	"net/http"

	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(cfg *config.Config, e *echo.Echo) {
	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})
}
