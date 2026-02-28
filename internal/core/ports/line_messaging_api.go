package ports

import (
	"context"

	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
)

type LineMessagingAPI interface {
	Reply(ctx context.Context, replyToken string, messages []domain.Message) error
}
