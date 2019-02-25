package api

import (
	"errors"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/matrix/xgbkb/business"
	"gitlab.wallstcn.com/matrix/xgbkb/types"
)

func ApiListChains(ctx echo.Context) (interface{}, error) {
	page, _ := strconv.ParseInt(ctx.Param("page"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Param("limit"), 10, 64)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		return nil, errors.New("Invalid limit")
	}
	return business.ListChains(page, limit)
}

func ApiCreateChain(ctx echo.Context) (interface{}, error) {
	var chainIn types.ChainIn
	if err := ctx.Bind(&chainIn); err != nil {
		return nil, err
	}
	if strings.TrimSpace(chainIn.Name) == "" {
		return nil, errors.New("Chain name empty")
	}
	if !business.IsValidChainName(chainIn.Name) {
		return nil, errors.New("Invalid chain name")
	}
	return business.CreateChain(&chainIn)
}

func ApiGetChain(ctx echo.Context) (interface{}, error) {
	name := strings.TrimSpace(ctx.QueryParam("name"))
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("Chain name empty")
	}
	if !business.IsValidChainName(name) {
		return nil, errors.New("Invalid chain name")
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	return business.GetChain(id, name)
}

func ApiUpdateChain(ctx echo.Context) (interface{}, error) {
	var updateChainIn types.UpdateChainIn
	if err := ctx.Bind(&updateChainIn); err != nil {
		return nil, err
	}
	chainId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	updateChainIn.Id = chainId
	if !business.IsValidChainName(updateChainIn.Name) {
		return nil, errors.New("New chain name invalid")
	}
	return business.UpdateChain(&updateChainIn)
}

func ApiAddProductToChain(ctx echo.Context) (interface{}, error) {
	var addProductToChainIn types.AddProductToChainIn
	if err := ctx.Bind(&addProductToChainIn); err != nil {
		return nil, err
	}
	if strings.TrimSpace(addProductToChainIn.ChainName) == "" {
		return nil, errors.New("Chain name empty")
	}
	if !business.IsValidChainName(addProductToChainIn.ChainName) {
		return nil, errors.New("Invalid chain name")
	}
	if strings.TrimSpace(addProductToChainIn.ProductName) == "" {
		return nil, errors.New("Product name empty")
	}
	return business.AddProductToChain(&addProductToChainIn)
}

func ApiRemoveProductFromChain(ctx echo.Context) (interface{}, error) {
	var remProductFromChainIn types.RemProductFromChainIn
	if err := ctx.Bind(&remProductFromChainIn); err != nil {
		return nil, err
	}
	return business.RemProductFromChain(&remProductFromChainIn)
}

func ApiGetProductsOfChain(ctx echo.Context) (interface{}, error) {
	var getProductsOfChainIn types.GetProductsOfChainIn
	if err := ctx.Bind(&getProductsOfChainIn); err != nil {
		return nil, err
	}
	if strings.TrimSpace(getProductsOfChainIn.ChainName) == "" {
		return nil, errors.New("Chain name empty")
	}
	if !business.IsValidChainName(getProductsOfChainIn.ChainName) {
		return nil, errors.New("Invalid chain name")
	}
	return business.GetProductsOfChain(getProductsOfChainIn.ChainId, getProductsOfChainIn.ChainName)
}
