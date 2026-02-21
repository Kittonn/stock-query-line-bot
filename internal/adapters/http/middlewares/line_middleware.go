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

			if !validateSignature(channelSecret, body, signature) {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid signature")
			}

			return next(c)
		}
	}
}

func validateSignature(channelSecret string, body []byte, signature string) bool {
	mac := hmac.New(sha256.New, []byte(channelSecret))
	mac.Write(body)
	expected := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(expected), []byte(signature))
}
