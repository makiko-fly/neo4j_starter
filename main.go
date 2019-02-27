package main

import (
	"fmt"

	"github.com/labstack/echo"
	echoware "github.com/labstack/echo/middleware"
	"gitlab.wallstcn.com/matrix/xgbkb/api"
	"gitlab.wallstcn.com/matrix/xgbkb/business"
	"gitlab.wallstcn.com/matrix/xgbkb/g"
	"gitlab.wallstcn.com/matrix/xgbkb/middleware"
	"gitlab.wallstcn.com/matrix/xgbkb/schedule"
	"gitlab.wallstcn.com/matrix/xgbkb/std"
	"gitlab.wallstcn.com/matrix/xgbkb/std/redislogger"
)

func main() {
	if err := std.LoadConf(&g.SysConf, "conf/xgbkb.yaml"); err != nil {
		panic(err)
	}
	redislogger.Init(g.SysConf.RedisLogger.Host, g.SysConf.RedisLogger.Port, g.SysConf.RedisLogger.Auth)
	redislogger.Printf("Loaded conf: \n%v", g.SysConf)

	g.InitShanghaiTimeZone()
	g.InitRedisClients()
	g.InitDb()

	business.InitNeo4j()

	schedule.StartJobs()
	defer schedule.StopJobs()

	e := echo.New()
	e.Use(middleware.LogRequest)
	e.Use(echoware.CORSWithConfig(echoware.DefaultCORSConfig))

	baseGroup := e.Group("")
	api.RegisterHttpPaths(baseGroup)
	err := e.Start(fmt.Sprintf("0.0.0.0:%d", g.SysConf.Http.Port))
	if err != nil {
		redislogger.Errf("Main, http server fails to start, err: %v", err)
	}
}
