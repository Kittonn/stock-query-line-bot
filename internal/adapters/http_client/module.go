package http_client

import "go.uber.org/fx"

var Module = fx.Module("http_client",
	fx.Provide(NewHTTPClient),
)
