package handlers

type Handlers struct {
	Health      *HealthHandler
	LineWebhook *LineWebhookHandler
}

func NewHandlers(
	health *HealthHandler,
	lineWebhook *LineWebhookHandler,
) *Handlers {
	return &Handlers{
		Health:      health,
		LineWebhook: lineWebhook,
	}
}
