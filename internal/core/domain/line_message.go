package domain

type Message interface{}

type TextMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type FlexMessage struct {
	Type     string `json:"type"`
	AltText  string `json:"altText"`
	Contents any    `json:"contents"`
}

func NewTextMessage(text string) TextMessage {
	return TextMessage{Type: "text", Text: text}
}

func NewFlexMessage(altText string, contents any) FlexMessage {
	return FlexMessage{Type: "flex", AltText: altText, Contents: contents}
}
