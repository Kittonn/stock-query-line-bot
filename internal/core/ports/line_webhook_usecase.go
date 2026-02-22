package ports

import (
	"context"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
)

type LineWebhookUsecase interface {
	HandleEvent(ctx context.Context, event *domain.LineEvent)
}
