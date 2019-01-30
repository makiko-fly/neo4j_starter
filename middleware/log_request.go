package middleware

import (
	"github.com/labstack/echo"
	"gitlab.wallstcn.com/matrix/xgbkb/std/logger"
)

func LogRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		logger.Infoln("=========================== New Request Received ===========================")
		logger.Infof("=== %s %s", r.Method, r.RequestURI)
		logger.Infof("Request: %v", r)
		return next(c)
	}
}
