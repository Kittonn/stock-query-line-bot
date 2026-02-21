package finnhubapi

import "go.uber.org/fx"

var Module = fx.Module("finnhubapi",
	fx.Provide(NewFinnhubAPI),
)
