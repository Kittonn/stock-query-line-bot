package http

import "go.uber.org/fx"

var Module = fx.Module("http",
	fx.Provide(
		NewEcho,
		NewHTTPServer,
	),
	fx.Invoke(
		RegisterRoutes,
		RunHTTPServer,
	),
)
