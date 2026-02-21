package domain

import "fmt"

type LineEventType string

const (
	LineEventTypeMessage  LineEventType = "message"
	LineEventTypePostback LineEventType = "postback"
)

type LineEvent struct {
	Type      LineEventType
	UserID    string
	Message   string
	Timestamp int64
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
