package main

import (
	"fmt"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/api"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/g"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/middleware"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/std"
	"gitlab.wallstcn.com/baoer/matrix/xgbkb/std/redislogger"
)

func main() {
	if err := std.LoadConf(&g.SysConf, "conf/xgbkb.yaml"); err != nil {
		panic(err)
	}
	redislogger.Init(g.SysConf.RedisLogger.Host, g.SysConf.RedisLogger.Port, g.SysConf.RedisLogger.Auth)
	redislogger.Printf("Loaded conf: \n%v", g.SysConf)

	e := echo.New()
	e.Use(middleware.LogRequest)
	baseGroup := e.Group("")
	api.RegisterHttpPaths(baseGroup)
	err := e.Start(fmt.Sprintf("127.0.0.1:%d", g.SysConf.Http.Port))
	if err != nil {
		redislogger.Errf("Main, http server fails to start, err: %v", err)
	}
}
