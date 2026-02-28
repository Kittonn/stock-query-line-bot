package line_messaging_api

import (
	"context"

	"fmt"

	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"github.com/Kittonn/stock-query-line-bot/internal/core/domain"
	"github.com/Kittonn/stock-query-line-bot/internal/core/ports"
	"github.com/go-resty/resty/v2"
)

type LineMessagingAPI struct {
	cfg   *config.Config
	resty *resty.Client
}

type ReplyMessageResponse struct {
	SentMessages []SentMessage `json:"sentMessages,omitempty"`
}

type SentMessage struct {
	ID         string `json:"id"`
	QuoteToken string `json:"quoteToken,omitempty"`
}

type ReplyMessageRequest struct {
	ReplyToken string `json:"replyToken"`
	Messages   []any  `json:"messages"`
}

func NewLineMessagingAPI(cfg *config.Config, client *resty.Client) ports.LineMessagingAPI {
	return &LineMessagingAPI{
		cfg:   cfg,
		resty: client,
	}
}

func (l *LineMessagingAPI) Reply(ctx context.Context, replyToken string, messages []domain.Message) error {
	msgs := make([]any, len(messages))
	for i, msg := range messages {
		msgs[i] = msg
	}

	replyMessageRequest := &ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages:   msgs,
	}

	resp, err := l.resty.R().SetContext(ctx).
		SetHeader("Authorization", "Bearer "+l.cfg.LineChannelAccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(replyMessageRequest).
		SetResult(&ReplyMessageResponse{}).
		Post(l.cfg.LineMessagingAPIURL + "/v2/bot/message/reply")

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("line api error: %s body: %s", resp.Status(), resp.String())
	}

	return nil
}
