package api

import (
	"errors"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/business"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/types"
)

func ApiListChains(ctx echo.Context) (interface{}, error) {
	var listChainsIn types.ListChainsIn
	if err := ctx.Bind(&listChainsIn); err != nil {
		return nil, err
	}
	if listChainsIn.Page <= 0 {
		listChainsIn.Page = 1
	}
	if listChainsIn.Limit <= 0 {
		listChainsIn.Limit = 10
	}
	if listChainsIn.Limit > 50 {
		return nil, errors.New("Invalid limit")
	}
	return business.ListChains(listChainsIn.Page, listChainsIn.Limit)
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
