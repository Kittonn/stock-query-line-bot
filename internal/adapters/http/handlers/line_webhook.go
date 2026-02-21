package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

type LineWebhookHandler struct{}

func NewLineWebhookHandler() *LineWebhookHandler {
	return &LineWebhookHandler{}
}

func (h *LineWebhookHandler) Handle(c *echo.Context) error {
	fmt.Println("Body", c.Request().Body)

	return c.NoContent(http.StatusNoContent)
}
