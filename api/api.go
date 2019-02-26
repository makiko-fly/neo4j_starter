package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/matrix/xgbkb/std"
	"gitlab.wallstcn.com/matrix/xgbkb/std/redislogger"
)

func RegisterHttpPaths(g *echo.Group) {
	registerHealthCheck(g)
	registerRedislog(g)
	registerAdminApis(g)
}

func registerHealthCheck(g *echo.Group) {
	g.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "You Got It!")
	})
}

func registerRedislog(g *echo.Group) {
	redislogger.MountRedisLog(g)
}

func registerAdminApis(g *echo.Group) {
	adminGroup := g.Group("/admin")
	adminGroup.GET("/search/byName", WrapRespAsJson(ApiSearchByName))
	adminGroup.POST("/product", WrapRespAsJson(ApiCreateProduct))
	adminGroup.PUT("/products/:id", WrapRespAsJson(ApiUpdateProduct))
	adminGroup.DELETE("/products/:id", WrapRespAsJson(ApiDeleteProduct))
	adminGroup.GET("/node/directlyRelated", WrapRespAsJson(ApiGetDirectlyRelatedNodes))
	adminGroup.POST("/relation", WrapRespAsJson(ApiCreateRelation))
	adminGroup.DELETE("/relation/:id", WrapRespAsJson(ApiDeleteRelation))
	adminGroup.GET("/chain/:id", WrapRespAsJson(ApiGetChain))
	adminGroup.GET("/chains/list", WrapRespAsJson(ApiListChains))
	adminGroup.POST("/chain", WrapRespAsJson(ApiCreateChain))
	adminGroup.PUT("/chain/:id", WrapRespAsJson(ApiUpdateChain))
	adminGroup.DELETE("/chain/:id", WrapRespAsJson(ApiDeleteChain))
	adminGroup.GET("/chain/products", WrapRespAsJson(ApiGetProductsOfChain))
	adminGroup.POST("/chain/addProduct", WrapRespAsJson(ApiAddProductToChain))
	adminGroup.POST("/chain/remProduct", WrapRespAsJson(ApiRemoveProductFromChain))
}

// ==================================================================

type Response struct {
	Code       int64           `json:"code"`
	Msg        string          `json:"msg"`
	DisplayMsg string          `json:"display_msg"`
	Data       json.RawMessage `json:"data"`
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
	var isAdminApi = false
	if strings.Index(ctx.Path(), "/admin/") == 0 {
		isAdminApi = true
	}

	resp := new(Response)
	resp.Data = []byte("{}")
	// set code
	stdErr, ok := err.(*std.Err)
	if ok {
		resp.Code = stdErr.Code
	} else {
		resp.Code = std.DefaultErrCode
	}
	// set message
	if isAdminApi {
		resp.DisplayMsg = err.Error()
	} else {
		resp.DisplayMsg = "something went wrong"
	}

	return ctx.JSON(http.StatusOK, resp)
}

func wrapNormalRespAsJson(ctx echo.Context, data interface{}) error {
	resp := new(Response)
	resp.Code = std.SuccessCode
	resp.Msg = ""
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
			resp.Msg = "fails to marshal response as json"
			resp.Data = []byte("{}")
		} else {
			resp.Data = dataBytes
		}
	}
	return ctx.JSON(http.StatusOK, resp)
}
