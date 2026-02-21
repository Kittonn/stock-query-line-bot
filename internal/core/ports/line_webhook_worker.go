package ports

import (
	"context"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
)

type LineWebhookWorker interface {
	Start(ctx context.Context, workerCount int)
	Enqueue(event *domain.LineEvent)
	Stop()
}
