package http

import (
	"github.com/Kittonn/stock-query-line-bot/internal/adapters/http/handlers"
	"go.uber.org/fx"
)

var Module = fx.Module("http",
	handlers.Module,
	fx.Provide(
		NewEcho,
		NewHTTPServer,
	),
	fx.Invoke(
		RegisterRoutes,
		RunHTTPServer,
	),
)
