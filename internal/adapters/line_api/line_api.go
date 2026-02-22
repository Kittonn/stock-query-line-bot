package line_api

import (
	"context"
	"encoding/json"

	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"github.com/go-resty/resty/v2"
)

type LineAPI struct {
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
	ReplyToken string            `json:"replyToken"`
	Messages   []json.RawMessage `json:"messages"`
}

func NewLineAPI(cfg *config.Config, client *resty.Client) *LineAPI {
	return &LineAPI{
		cfg:   cfg,
		resty: client,
	}
}

func (l *LineAPI) Reply(ctx context.Context, replyMessageRequest *ReplyMessageRequest) (*ReplyMessageResponse, error) {
	resp, err := l.resty.R().SetContext(ctx).
		SetHeader("Authorization", "Bearer "+l.cfg.LineChannelAccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(replyMessageRequest).
		SetResult(&ReplyMessageResponse{}).
		Post(l.cfg.LineMessagingAPIURL + "/v2/bot/message/reply")

	if err != nil {
		return nil, err
	}

	return resp.Result().(*ReplyMessageResponse), nil
}
