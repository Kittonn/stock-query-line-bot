package usecases

import (
	"context"

	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
	"go.uber.org/fx"
)

var Module = fx.Module("usecases", fx.Provide(
	NewLineWebhookWorker,
	NewLineWebhookUsecase,
	NewStockUsecase,
),
	fx.Invoke(StartLineWorker),
)

func StartLineWorker(lc fx.Lifecycle, worker ports.LineWebhookWorker) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			worker.Start(ctx, 5)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			worker.Stop()
			return nil
		},
	})
}
