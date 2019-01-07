package api

import (
	"github.com/labstack/echo"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/business"
)

func SearchByName(ctx echo.Context) (interface{}, error) {
	return business.SearchAllWithNameLikeKeywoard("key")
}
