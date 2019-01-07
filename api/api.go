package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/std"
)

func RegisterHttpPaths(g *echo.Group) {
	registerHealthCheck(g)
	registerAdminApis(g)
}

func registerHealthCheck(g *echo.Group) {
	g.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "You Got It!")
	})
}

func registerAdminApis(g *echo.Group) {
	adminGroup := g.Group("/admin")
	adminGroup.GET("/search/byName/:keyword", WrapRespAsJson(SearchByName))
}

// ==================================================================

type Response struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type (
	CommonHttpHandler func(ctx echo.Context) (interface{}, error)
)

func WrapRespAsJson(handler CommonHttpHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		value, err := handler(ctx)
		if err != nil {
			return wrapErrRespAsJson(ctx, err)
		} else {
			return wrapNormalRespAsJson(ctx, value)
		}
	}
}

func wrapErrRespAsJson(ctx echo.Context, err error) error {
	resp := new(Response)
	resp.Data = []byte("{}")
	errWithCode, ok := err.(*std.ErrWithCode)
	if ok {
		resp.Code = errWithCode.Code
		resp.Message = errWithCode.Error()
	} else {
		resp.Code = std.DefaultErrCode
		resp.Message = err.Error()
	}
	return ctx.JSON(http.StatusOK, resp)
}

func wrapNormalRespAsJson(ctx echo.Context, data interface{}) error {
	resp := new(Response)
	resp.Code = std.SuccessCode
	resp.Message = ""
	// assemble data
	if data == nil {
		data = make(map[string]interface{})
	}
	if byteArr, ok := data.([]byte); ok {
		// don't marshal
		resp.Data = byteArr
	} else {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			// change return code
			resp.Code = std.DefaultErrCode
			resp.Message = "fails to marshal response as json"
			resp.Data = []byte("{}")
		} else {
			resp.Data = dataBytes
		}
	}
	return ctx.JSON(http.StatusOK, resp)
}
