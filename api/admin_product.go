package api

import (
	"errors"
	"strings"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/business"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/types"
)

func CreateProduct(ctx echo.Context) (interface{}, error) {
	var productIn types.ProductIn
	if err := ctx.Bind(&productIn); err != nil {
		return nil, err
	}
	if strings.TrimSpace(productIn.Name) == "" {
		return nil, errors.New("Product name is empty")
	}
	if !IsValidProductName(productIn.Name) {
		return nil, errors.New("Product name is invalid")
	}
	return business.CreateProduct(&productIn)
}

func IsValidProductName(name string) bool {
	return true
}

func UpdateProduct(ctx echo.Context) (interface{}, error) {
	var productIn types.ProductIn
	if err := ctx.Bind(&productIn); err != nil {
		return nil, err
	}
	oldName := strings.TrimSpace(ctx.Param("oldName"))
	if oldName == "" {
		return nil, errors.New("Product's old name in path is empty")
	}
	if strings.TrimSpace(productIn.Name) == "" {
		return nil, errors.New("New product name is empty")
	}
	if !IsValidProductName(productIn.Name) {
		return nil, errors.New("Product name is invalid")
	}
	return business.UpdateProduct(oldName, &productIn)
}
