package logger

import (
	"fmt"

	"gitlab.wallstcn.com/matrix/xgbkb/std/redislogger"
)

func Infoln(args ...interface{}) {
	logLine := fmt.Sprint(args...)
	var trackId int32 = 0
	wrappedLogLine := redislogger.WrapLogLine(trackId, logLine)
	fmt.Println(wrappedLogLine)
}

func Infof(format string, args ...interface{}) {
	logLine := fmt.Sprintf(format, args...)
	var trackId int32 = 0
	wrappedLogLine := redislogger.WrapLogLine(trackId, logLine)
	fmt.Println(wrappedLogLine)
}
