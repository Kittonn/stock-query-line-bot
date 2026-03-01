package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"github.com/Kittonn/stock-query-line-bot/pkg/logger"
	"github.com/labstack/echo/v5"
	"go.uber.org/fx"
)

func NewHTTPServer(e *echo.Echo, cfg *config.Config) *http.Server {
	return &http.Server{
		Addr:              ":" + fmt.Sprint(cfg.Port),
		Handler:           e,
		ReadHeaderTimeout: 10 * time.Second,
	}
}

func RunHTTPServer(lc fx.Lifecycle, srv *http.Server, log logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					panic(err)
				}
			}()

			log.Info("Starting HTTP server on port ", srv.Addr)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			return srv.Shutdown(shutdownCtx)
		},
	})
}
