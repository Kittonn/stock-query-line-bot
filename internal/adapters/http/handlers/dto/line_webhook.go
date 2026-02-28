package dto

import "github.com/Kittonn/stock-query-line-bot/internal/core/domain"

type WebhookRequest struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}

type Event struct {
	Type            string           `json:"type"`
	Timestamp       int64            `json:"timestamp"`
	Source          Source           `json:"source,omitempty"`
	WebhookEventId  string           `json:"webhookEventId"`
	Mode            EventMode        `json:"mode"`
	ReplyToken      string           `json:"replyToken,omitempty"`
	DeliveryContext DeliveryContext  `json:"deliveryContext"`
	Message         *MessageContent  `json:"message,omitempty"`
	Postback        *PostbackContent `json:"postback,omitempty"`
}

type MessageContent struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Text string `json:"text,omitempty"`
}

type EventMode string

type Source struct {
	UserId string `json:"userId"`
	Type   string `json:"type"`
}

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

type PostbackContent struct {
	Data   string          `json:"data"`
	Params *PostbackParams `json:"params,omitempty"`
}

type PostbackParams struct {
	Date     string `json:"date,omitempty"`
	Time     string `json:"time,omitempty"`
	Datetime string `json:"datetime,omitempty"`
}

func (e *Event) ToDomain() (*domain.LineEvent, error) {
	t, err := domain.ParseLineEventType(e.Type)
	if err != nil {
		return nil, err
	}

	event := &domain.LineEvent{
		Type:       t,
		UserID:     e.Source.UserId,
		Timestamp:  e.Timestamp,
		ReplyToken: e.ReplyToken,
	}

	if e.Message != nil {
		event.Message = &domain.MessagePayload{
			Text: e.Message.Text,
		}
	}

	if e.Postback != nil {
		event.Postback = &domain.PostbackPayload{
			Data: e.Postback.Data,
		}
	}

	return event, nil
}
