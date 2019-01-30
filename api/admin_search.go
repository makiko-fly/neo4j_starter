package api

import (
	"errors"
	"strings"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/matrix/xgbkb/business"
)

func ApiSearchByName(ctx echo.Context) (interface{}, error) {
	keywordStr := strings.TrimSpace(ctx.QueryParam("keyword"))
	if len(keywordStr) == 0 {
		return nil, errors.New("empty keyword")
	}
	return business.SearchAllWithNameLikeKeywoard(keywordStr)
}
