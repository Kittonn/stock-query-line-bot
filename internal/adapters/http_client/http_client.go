package http_client

import "github.com/go-resty/resty/v2"

func NewHTTPClient() *resty.Client {
	return resty.New()
}
