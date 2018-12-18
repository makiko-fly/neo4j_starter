package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterHttpPaths(g *echo.Group) {
	registerHealthCheck(g)
}

func registerHealthCheck(g *echo.Group) {
	g.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "You Got It!")
	})
}
