package domain

import "fmt"

type LineEventType string

const (
	LineEventTypeMessage  LineEventType = "message"
	LineEventTypePostback LineEventType = "postback"
)

type MessagePayload struct {
	Text string
}

type PostbackPayload struct {
	Data string
}

type LineEvent struct {
	Type       LineEventType
	UserID     string
	Timestamp  int64
	ReplyToken string

	Message  *MessagePayload
	Postback *PostbackPayload
}

func ParseLineEventType(s string) (LineEventType, error) {
	switch s {
	case "message":
		return LineEventTypeMessage, nil
	case "postback":
		return LineEventTypePostback, nil
	default:
		return "", fmt.Errorf("unknown event type: %s", s)
	}
}
