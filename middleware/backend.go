package middleware

import (
	"github.com/STNS/STNS/model"
	"github.com/labstack/echo"
)

const (
	BackendKey = "Backend"
)

func Backend(b model.Backend) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			c.Set(BackendKey, b)
			return next(c)
		})
	}
}
