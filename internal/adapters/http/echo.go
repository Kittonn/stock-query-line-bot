package http

import "github.com/labstack/echo/v5"

func NewEcho() *echo.Echo {
	e := echo.New()

	return e
}
