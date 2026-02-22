package finnhub_api

import "go.uber.org/fx"

var Module = fx.Module("finnhubapi",
	fx.Provide(NewFinnhubAPI),
)
