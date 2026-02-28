package line_messaging_api

import "go.uber.org/fx"

var Module = fx.Module("lineapi",
	fx.Provide(NewLineMessagingAPI),
)
