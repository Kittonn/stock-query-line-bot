package middlewares

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/labstack/echo/v5"
)

const LineSignatureHeader = "X-Line-Signature"

func VerifyLineSignature(channelSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			req := c.Request()

			body, err := io.ReadAll(req.Body)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "cannot read body")
			}

			req.Body = io.NopCloser(io.MultiReader(bytes.NewReader(body), req.Body))

			signature := req.Header.Get(LineSignatureHeader)
			if signature == "" {
				return echo.NewHTTPError(http.StatusBadRequest, "missing X-Line-Signature header")
			}

			if !ValidateSignature(channelSecret, signature, body) {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid signature")
			}

			return next(c)
		}
	}
}

// Reference: https://github.com/line/line-bot-sdk-go/blob/master/linebot/webhook/parse.go
func ValidateSignature(channelSecret, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(channelSecret))

	_, err = hash.Write(body)
	if err != nil {
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}
