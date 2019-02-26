package api

import (
	"errors"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/matrix/xgbkb/business"
	"gitlab.wallstcn.com/matrix/xgbkb/g"
	"gitlab.wallstcn.com/matrix/xgbkb/types"
)

func ApiCreateProduct(ctx echo.Context) (interface{}, error) {
	var productIn types.ProductIn
	if err := ctx.Bind(&productIn); err != nil {
		return nil, err
	}
	if strings.TrimSpace(productIn.Name) == "" {
		return nil, errors.New("Product name is empty")
	}
	if !business.IsValidProductName(productIn.Name) {
		return nil, errors.New("Product name is invalid")
	}
	return business.CreateProduct(&productIn)
}

func ApiUpdateProduct(ctx echo.Context) (interface{}, error) {
	var updateProductIn types.UpdateProductIn
	if err := ctx.Bind(&updateProductIn); err != nil {
		return nil, err
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	updateProductIn.Id = id
	if strings.TrimSpace(updateProductIn.Name) == "" {
		return nil, errors.New("New product name is empty")
	}
	if !business.IsValidProductName(updateProductIn.Name) {
		return nil, errors.New("New product name is invalid")
	}
	return business.UpdateProduct(id, updateProductIn.OldName, &updateProductIn.ProductIn)
}

func ApiDeleteProduct(ctx echo.Context) (interface{}, error) {
	var deleteProductIn types.DeleteProductIn
	if err := ctx.Bind(&deleteProductIn); err != nil {
		return nil, err
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	deleteProductIn.Id = id
	if strings.TrimSpace(deleteProductIn.Name) == "" {
		return nil, errors.New("Product name is empty")
	}
	if !business.IsValidProductName(deleteProductIn.Name) {
		return nil, errors.New("Product name is invalid")
	}
	return business.DeleteNode(id, deleteProductIn.Name, g.LabelProduct)
}
